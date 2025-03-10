package douyin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nanlei2000/douyin_download/pkg/model"
)

func (d *DouYin) GetAllVideoIDList(secUid string) ([]string, error) {
	init := true
	cursor := 0
	count := 35
	idList := []string{}

	for {
		if cursor == 0 && !init {
			break
		}

		apiURL := fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=%s&count=%d&max_cursor=%d&aid=1128&_signature=PDHVOQAAXMfFyj02QEpGaDwx1S&dytk=", secUid, count, cursor)
		header := http.Header{}
		header.Add("User-Agent", DefaultUserAgent)
		req, err := http.NewRequest(http.MethodGet, apiURL, nil)
		if err != nil {
			log.Printf("fail to get post, secUid: %s, err: %s", secUid, err)
			return []string{}, err
		}

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Printf("fail to get post, secUid: %s, err: %s", secUid, err)
			return []string{}, err
		}
		defer resp.Body.Close()

		d.printf("resp status: %s", resp.Status)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("fail to read body, secUid: %s, err: %s", secUid, err)
			return []string{}, err
		}

		var userPost model.UserPost
		err = json.Unmarshal(body, &userPost)

		if err != nil {
			log.Printf("fail to unmarshal json, secUid: %s, err: %s, body: %s", secUid, err, body)
			return []string{}, err
		}
		if userPost.StatusCode != 0 {
			return []string{}, fmt.Errorf("resp err, resp: %s", body)
		}

		init = false
		cursor = int(userPost.MaxCursor)
		for _, item := range userPost.AwemeList {
			idList = append(idList, item.AwemeID)
		}
	}

	return idList, nil
}
