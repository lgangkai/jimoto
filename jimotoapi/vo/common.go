package vo

type Image struct {
	Image string `json:"image"`
}

type UploadResp struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
}
