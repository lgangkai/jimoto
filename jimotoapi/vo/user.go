package vo

type User struct {
	UserId uint64 `json:"user_id"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
}

type CreateProfileReq struct {
	Username string `form:"username"`
	Birthday string `form:"birthday"`
	Email    string `form:"email"`
}

type AccountReq struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
