package douyin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nanlei2000/douyin_download/internal/utils"
)

type VideoType int

const (
	//VideoPlayType 视频类
	VideoPlayType VideoType = 0
	//ImagePlayType 图文类
	ImagePlayType VideoType = 1
)

type Video struct {
	VideoId         string   `json:"video_id"`
	PlayId          string   `json:"play_id"`
	PlayAddr        string   `json:"play_addr"`
	VideoRawAddr    string   `json:"video_raw_addr"`
	PlayRawAddr     string   `json:"play_raw_addr"`
	Cover           string   `json:"cover"`
	OriginCover     string   `json:"origin_cover"`
	OriginCoverList []string `json:"origin_cover_list"`
	Desc            string   `json:"desc"`
	RawLink         string   `json:"raw_link"`
	Author          struct {
		Id           string `json:"id"`
		ShortId      string `json:"short_id"`
		Nickname     string `json:"nickname"`
		AvatarLarger string `json:"avatar_larger"`
		Signature    string `json:"signature"`
	} `json:"author"`
	Images    []ImageItem `json:"images"`
	VideoType VideoType   `json:"video_type"`
}

type ImageItem struct {
	ImageUrl string `json:"image_url"`
	ImageId  string `json:"image_id"`
}

func (v *Video) GetFilename() string {
	if ext := filepath.Ext(v.PlayId); ext != "" {
		return v.VideoId + ext
	}
	return v.VideoId + ".mp4"
}

// Download 下载视频、图文到文件到指定目录，返回视频地址（图文为背景音乐视频地址）
func (v *Video) Download(distDir string) (path string, err error) {
	defer func() {
		if pErr := recover(); pErr != nil {
			log.Printf("出现panic: [filename=%s] [errmsg=%s]", distDir, err)
			err = fmt.Errorf("%s", pErr)
		}
	}()
	distDir, err = filepath.Abs(distDir)
	if err != nil {
		log.Printf("获取报错地址失败 [filename=%s] [error=%+v]", distDir, err)
		return "", err
	}
	folderName := fmt.Sprintf("%s_%s", v.Author.Nickname, v.Author.ShortId)
	distDir = filepath.Join(distDir, folderName, v.GetFilename())
	dir := filepath.Dir(distDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}
	//如果是图片类，则将图片下载到指定目录
	if v.VideoType == ImagePlayType {
		imagePath := filepath.Join(dir, v.VideoId)
		if err := os.MkdirAll(imagePath, os.ModePerm); err != nil {
			log.Printf("创建目录失败 [path=%s]", imagePath)
		}
		for _, image := range v.Images {
			ext := ".jpeg"
			uri, err := url.Parse(image.ImageUrl)
			if err != nil {
				log.Printf("解析图片地址失败 [image_url=%s] [errmsg=%+v]", image.ImageUrl, err)
			} else {
				ext = filepath.Ext(uri.Path)
			}
			imageId := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(image.ImageId, "//", ""), "\\\\", "/"), "/", "-")
			imageName := filepath.Join(imagePath, imageId+ext)

			if _, err := os.Stat(imageName); !os.IsNotExist(err) {
				log.Printf("图片本地已存在，跳过下载 [image_name=%s]", imageName)
				continue
			}

			req, err := http.NewRequest(http.MethodGet, image.ImageUrl, nil)
			if err != nil {
				log.Printf("下载图像出错 -> [play_id=%s] [image_url=%s] [errmsg=%+v]", v.PlayId, image.ImageUrl, err)
				continue
			}
			req.Header.Add("User-Agent", DefaultUserAgent)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("获取图像响应出错 -> [play_id=%s] [image_url=%s] [errmsg=%+v]", v.PlayId, image.ImageUrl, err)
				continue
			}

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("解析图像出错 -> [play_id=%s] [image_url=%s]", v.PlayId, image.ImageUrl)
				continue
			}
			_ = resp.Body.Close()
			err = ioutil.WriteFile(imageName, b, os.ModePerm)
			if err != nil {
				log.Printf("保存图像出错 -> [play_id=%s] [image_url=%s]", v.PlayId, image.ImageUrl)
				continue
			}

			log.Printf("图片数据 [image_url=%s] [image_name=%s]", image.ImageUrl, imageName)
			time.Sleep(time.Microsecond * 110)
		}
		//如果是图文，需要将音频和图像放入一个目录
		distDir = filepath.Join(imagePath, filepath.Base(distDir))
	}

	if _, err := os.Stat(distDir); !os.IsNotExist(err) {
		log.Printf("视频本地已存在，[filename=%s]", distDir)
		return distDir, nil
	}

	req, err := http.NewRequest(http.MethodGet, v.PlayAddr, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", DefaultUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	f1, err := os.Create(distDir)
	if err != nil {
		log.Printf("创建文件失败 [filename=%s] [errmsg=%+v]", distDir, err)
		return "", err
	}
	defer f1.Close()
	_, err = io.Copy(f1, resp.Body)
	if err != nil {
		log.Printf("创建文件失败 [filename=%s] [errmsg=%+v]", distDir, err)
		return "", err
	}

	log.Printf("写入文件成功： [filename=%s]", distDir)

	return distDir, nil
}

// DownloadCover 下载封面文件
func (v *Video) DownloadCover(urlStr string, filename string) (string, error) {
	uri, err := url.ParseRequestURI(urlStr)
	if err != nil {
		log.Printf("解析封面文件失败: url[%s] filename[%s] %+v", urlStr, filename, err)
		return "", err
	}

	filename = filepath.Join(filename, v.Author.Id, "cover", uri.Path)

	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0666); err != nil {
			return "", err
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("创建封面文件失败: url[%s] filename[%s] %+v", urlStr, filename, err)
		return "", err
	}
	defer utils.SafeClose(f)

	header := http.Header{}
	header.Add("User-Agent", DefaultUserAgent)
	header.Add("Upgrade-Insecure-Requests", "1")

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		log.Printf("下载封面文件失败: url[%s] filename[%s] %+v", urlStr, filename, err)
		return "", err
	}
	req.Header = header
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return "", err
	}
	defer utils.SafeClose(resp.Body)
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Printf("保存图片失败: %s  %+v", urlStr, err)
		return "", err
	}
	log.Printf("保存封面成功: %s  %s", urlStr, filename)
	return filename, nil
}

//GetDownloadUrl 获取下载链接
func (v *Video) GetDownloadUrl() (string, error) {
	req, err := http.NewRequest(http.MethodGet, v.PlayAddr, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", DefaultUserAgent)
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	lv := resp.Header.Get("Location")

	return lv, nil
}

func (v *Video) String() string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("编码失败 -> %s", err)
	} else {
		return string(b)
	}
	return fmt.Sprintf("%+v", *v)
}
