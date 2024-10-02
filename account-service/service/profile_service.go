package service

import (
	"context"
	errs "errs"
	"github.com/lgangkai/golog"
	"jimoto/account-service/dao"
	"jimoto/account-service/model"
)

type ProfileService struct {
	profileDao *dao.ProfileDao
	logger     *golog.Logger
}

func NewProfileService(profileDao *dao.ProfileDao, logger *golog.Logger) *ProfileService {
	return &ProfileService{
		profileDao: profileDao,
		logger:     logger,
	}
}

func (s *ProfileService) GetProfile(ctx context.Context, userId uint64) (*model.Profile, error) {
	s.logger.Info(ctx, "Call ProfileService.GetProfile")
	profile, err := s.profileDao.GetProfileByUserId(ctx, userId)
	if err != nil {
		s.logger.Error(ctx, "Fail to get profile, err:", err.Error())
		return nil, errs.New(errs.ERR_GET_PROFILE_FAILED)
	}
	return profile, nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	s.logger.Info(ctx, "Call ProfileService.UpdateProfile")
	err := s.profileDao.Update(ctx, profile)
	if err != nil {
		s.logger.Error(ctx, "Fail to update profile, err:", err.Error())
		return errs.New(errs.ERR_UPDATE_PROFILE_FAILED)
	}
	return nil
}

func (s *ProfileService) DeleteProfile(ctx context.Context, userId uint64) error {
	s.logger.Info(ctx, "Call ProfileService.DeleteProfile.")
	err := s.profileDao.Delete(ctx, userId)
	if err != nil {
		s.logger.Error(ctx, "Fail to delete profile, err:", err.Error())
		return errs.New(errs.ERR_DELETE_PROFILE_FAILED)
	}
	return nil
}

func (s *ProfileService) CreateProfile(ctx context.Context, profile *model.Profile) error {
	s.logger.Info(ctx, "Call ProfileService.CreateProfile, profile: ", profile)
	err := s.profileDao.Insert(ctx, profile)
	if err != nil {
		s.logger.Error(ctx, "Fail to delete profile, err:", err.Error())
		return errs.New(errs.ERR_CREATE_PROFILE_FAILED)
	}
	return nil
}
