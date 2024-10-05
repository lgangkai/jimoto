package model

const TAB_NAME_PROFILE = "profile_tab"

type Profile struct {
	Id        uint64 `json:"id"`
	UserId    uint64 `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	IsDeleted bool   `json:"is_deleted"`
}

func (p Profile) TableName() string {
	return TAB_NAME_PROFILE
}
