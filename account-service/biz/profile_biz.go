package biz

import (
	"context"
	"protos/account"
)

type ProfileBiz struct {
}

func NewProfileBiz() *ProfileBiz {
	return &ProfileBiz{}
}

func (b *ProfileBiz) GetProfile(context context.Context, in *account.GetProfileRequest, out *account.GetProfileResponse) error {
	return nil
}

func (b *ProfileBiz) DeleteProfile(traceContext context.Context, in *account.DeleteProfileRequest, out *account.DeleteProfileResponse) error {
	return nil
}

func (b *ProfileBiz) CreateProfile(traceContext context.Context, in *account.CreateProfileRequest, out *account.CreateProfileResponse) error {
	return nil
}

func (b *ProfileBiz) UpdateProfile(traceContext context.Context, in *account.UpdateProfileRequest, out *account.UpdateProfileResponse) error {
	return nil
}
