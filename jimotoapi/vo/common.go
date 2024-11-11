package vo

type Image struct {
	Image string `json:"image"`
}

type UploadResp struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
}

type AuthData struct {
	UserId uint64 `json:"user_Id"`
	Email  string `json:"mail"`
}
