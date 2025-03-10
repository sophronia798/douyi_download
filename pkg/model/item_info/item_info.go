package iteminfo

// Generated by https://quicktype.io

type ItemInfo struct {
	StatusCode int64         `json:"status_code"`
	ItemList   []ItemList    `json:"item_list"`
	FilterList []interface{} `json:"filter_list"`
	Extra      Extra         `json:"extra"`
}

type Extra struct {
	Now   int64  `json:"now"`
	Logid string `json:"logid"`
}

type ItemList struct {
	Geofencing   interface{}  `json:"geofencing"`
	Promotions   interface{}  `json:"promotions"`
	AnchorInfo   AnchorInfo   `json:"anchor_info"`
	ShareURL     string       `json:"share_url"`
	Duration     int64        `json:"duration"`
	CommentList  interface{}  `json:"comment_list"`
	IsLiveReplay bool         `json:"is_live_replay"`
	Author       Author       `json:"author"`
	TextExtra    []TextExtra  `json:"text_extra"`
	AuthorUserID int64        `json:"author_user_id"`
	LongVideo    interface{}  `json:"long_video"`
	IsPreview    int64        `json:"is_preview"`
	ForwardID    string       `json:"forward_id"`
	Desc         string       `json:"desc"`
	Images       []Image      `json:"images"`
	VideoText    interface{}  `json:"video_text"`
	Music        Music        `json:"music"`
	Video        Video        `json:"video"`
	ShareInfo    ShareInfo    `json:"share_info"`
	AwemeType    int64        `json:"aweme_type"`
	GroupID      int64        `json:"group_id"`
	CreateTime   int64        `json:"create_time"`
	RiskInfos    RiskInfos    `json:"risk_infos"`
	LabelTopText interface{}  `json:"label_top_text"`
	AwemePoiInfo AwemePoiInfo `json:"aweme_poi_info"`
	VideoLabels  interface{}  `json:"video_labels"`
	ChaList      []ChaList    `json:"cha_list"`
	Statistics   Statistics   `json:"statistics"`
	ImageInfos   []ImageInfo  `json:"image_infos"`
	GroupIDStr   string       `json:"group_id_str"`
	AwemeID      string       `json:"aweme_id"`
}

type AnchorInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type int64  `json:"type"`
}

type Author struct {
	PlatformSyncInfo interface{}  `json:"platform_sync_info"`
	Geofencing       interface{}  `json:"geofencing"`
	PolicyVersion    interface{}  `json:"policy_version"`
	ShortID          string       `json:"short_id"`
	AvatarMedium     AvatarLarger `json:"avatar_medium"`
	FollowStatus     int64        `json:"follow_status"`
	FollowersDetail  interface{}  `json:"followers_detail"`
	TypeLabel        interface{}  `json:"type_label"`
	Signature        string       `json:"signature"`
	UniqueID         string       `json:"unique_id"`
	CardEntries      interface{}  `json:"card_entries"`
	MixInfo          interface{}  `json:"mix_info"`
	Uid              string       `json:"uid"`
	AvatarLarger     AvatarLarger `json:"avatar_larger"`
	AvatarThumb      AvatarLarger `json:"avatar_thumb"`
	Nickname         string       `json:"nickname"`
}

type AvatarLarger struct {
	URI     string   `json:"uri"`
	URLList []string `json:"url_list"`
}

type AwemePoiInfo struct {
	Tag      string       `json:"tag"`
	Icon     AvatarLarger `json:"icon"`
	PoiName  string       `json:"poi_name"`
	TypeName string       `json:"type_name"`
}

type ChaList struct {
	UserCount      int64        `json:"user_count"`
	ConnectMusic   interface{}  `json:"connect_music"`
	Type           int64        `json:"type"`
	Cid            string       `json:"cid"`
	Desc           string       `json:"desc"`
	ViewCount      int64        `json:"view_count"`
	HashTagProfile string       `json:"hash_tag_profile"`
	IsCommerce     bool         `json:"is_commerce"`
	ChaName        string       `json:"cha_name"`
	CoverItem      AvatarLarger `json:"cover_item"`
}

type ImageInfo struct {
}

type Image struct {
	URI             string   `json:"uri"`
	URLList         []string `json:"url_list"`
	DownloadURLList []string `json:"download_url_list"`
	Height          int64    `json:"height"`
	Width           int64    `json:"width"`
}

type Music struct {
	Author      string       `json:"author"`
	CoverHD     AvatarLarger `json:"cover_hd"`
	CoverLarge  AvatarLarger `json:"cover_large"`
	PlayURL     AvatarLarger `json:"play_url"`
	Duration    int64        `json:"duration"`
	Status      int64        `json:"status"`
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	CoverThumb  AvatarLarger `json:"cover_thumb"`
	Position    interface{}  `json:"position"`
	Mid         string       `json:"mid"`
	CoverMedium AvatarLarger `json:"cover_medium"`
}

type RiskInfos struct {
	Warn             bool   `json:"warn"`
	Type             int64  `json:"type"`
	Content          string `json:"content"`
	ReflowUnplayable int64  `json:"reflow_unplayable"`
}

type ShareInfo struct {
	ShareDesc      string `json:"share_desc"`
	ShareTitle     string `json:"share_title"`
	ShareWeiboDesc string `json:"share_weibo_desc"`
}

type Statistics struct {
	DiggCount    int64  `json:"digg_count"`
	PlayCount    int64  `json:"play_count"`
	ShareCount   int64  `json:"share_count"`
	AwemeID      string `json:"aweme_id"`
	CommentCount int64  `json:"comment_count"`
}

type TextExtra struct {
	Start       int64  `json:"start"`
	End         int64  `json:"end"`
	Type        int64  `json:"type"`
	HashtagName string `json:"hashtag_name"`
	HashtagID   int64  `json:"hashtag_id"`
}

type Video struct {
	Height       int64        `json:"height"`
	OriginCover  AvatarLarger `json:"origin_cover"`
	Duration     int64        `json:"duration"`
	Vid          string       `json:"vid"`
	PlayAddr     AvatarLarger `json:"play_addr"`
	Cover        AvatarLarger `json:"cover"`
	Width        int64        `json:"width"`
	Ratio        string       `json:"ratio"`
	HasWatermark bool         `json:"has_watermark"`
	BitRate      interface{}  `json:"bit_rate"`
}
