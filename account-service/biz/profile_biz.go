package biz

import (
	"context"
	"github.com/lgangkai/golog"
	"jimoto/account-service/model"
	"jimoto/account-service/service"
	"protos/account"
)

type ProfileBiz struct {
	profileService *service.ProfileService
	logger         *golog.Logger
}

func NewProfileBiz(profileService *service.ProfileService, logger *golog.Logger) *ProfileBiz {
	return &ProfileBiz{
		profileService: profileService,
		logger:         logger,
	}
}

func (b *ProfileBiz) GetProfile(context context.Context, in *account.GetProfileRequest, out *account.GetProfileResponse) error {
	b.logger.Info(context, "Call ProfileBiz.GetProfile, request: ", in)
	id := in.GetUserId()
	p, err := b.profileService.GetProfile(context, id)
	if err != nil {
		b.logger.Error(context, "Get profile failed, err: ", err.Error())
		return err
	}
	out.Profile = &account.Profile{
		Id:       p.Id,
		UserId:   p.UserId,
		Username: p.Username,
		Email:    p.Email,
		Avatar:   p.AvatarUrl,
	}
	b.logger.Info(context, "Call ProfileBiz.GetProfile successfully.")
	return nil
}

func (b *ProfileBiz) DeleteProfile(traceContext context.Context, in *account.DeleteProfileRequest, out *account.DeleteProfileResponse) error {
	b.logger.Info(traceContext, "Call ProfileBiz.DeleteProfile, request: ", in)
	id := in.GetUserId()
	err := b.profileService.DeleteProfile(traceContext, id)
	if err != nil {
		b.logger.Error(traceContext, "Delete profile failed, err: ", err.Error())
		return err
	}
	b.logger.Info(traceContext, "Call ProfileBiz.DeleteProfile successfully.")
	return nil
}

func (b *ProfileBiz) CreateProfile(traceContext context.Context, in *account.CreateProfileRequest, out *account.CreateProfileResponse) error {
	b.logger.Info(traceContext, "Call ProfileBiz.CreateProfile, request: ", in)
	p := in.GetProfile()
	mp := &model.Profile{
		Id:        p.Id,
		UserId:    p.UserId,
		Username:  p.Username,
		Email:     p.Email,
		AvatarUrl: p.Avatar,
	}
	err := b.profileService.CreateProfile(traceContext, mp)
	if err != nil {
		b.logger.Error(traceContext, "Create profile failed, err: ", err.Error())
		return err
	}
	b.logger.Info(traceContext, "Call ProfileBiz.CreateProfile successfully.")
	return nil
}

func (b *ProfileBiz) UpdateProfile(traceContext context.Context, in *account.UpdateProfileRequest, out *account.UpdateProfileResponse) error {
	b.logger.Info(traceContext, "Call ProfileBiz.UpdateProfile, request: ", in)
	p := in.GetProfile()
	mp := &model.Profile{
		Id:        p.Id,
		UserId:    p.UserId,
		Username:  p.Username,
		Email:     p.Email,
		AvatarUrl: p.Avatar,
	}
	err := b.profileService.UpdateProfile(traceContext, mp)
	if err != nil {
		b.logger.Error(traceContext, "Update profile failed, err: ", err.Error())
		return err
	}
	b.logger.Info(traceContext, "Call ProfileBiz.UpdateProfile successfully.")
	return nil
}
