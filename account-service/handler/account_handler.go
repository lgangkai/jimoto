package handler

import (
	"context"
	"jimoto/account-service/biz"
	"protos/account"
)

type AccountHandlerImpl struct {
	accountBiz *biz.AccountBiz
	profileBiz *biz.ProfileBiz
}

func NewAccountHandlerImpl(accountBiz *biz.AccountBiz, profileBiz *biz.ProfileBiz) *AccountHandlerImpl {
	return &AccountHandlerImpl{accountBiz: accountBiz, profileBiz: profileBiz}
}

func (h *AccountHandlerImpl) GetProfile(ctx context.Context, in *account.GetProfileRequest, out *account.GetProfileResponse) error {
	return h.profileBiz.GetProfile(getTraceContext(ctx, in.GetRequestId(), in.GetUserId()), in, out)
}

func (h *AccountHandlerImpl) DeleteProfile(ctx context.Context, in *account.DeleteProfileRequest, out *account.DeleteProfileResponse) error {
	return h.profileBiz.DeleteProfile(getTraceContext(ctx, in.GetRequestId(), in.GetUserId()), in, out)
}

func (h *AccountHandlerImpl) CreateProfile(ctx context.Context, in *account.CreateProfileRequest, out *account.CreateProfileResponse) error {
	return h.profileBiz.CreateProfile(getTraceContext(ctx, in.GetRequestId(), in.GetProfile().GetUserId()), in, out)
}

func (h *AccountHandlerImpl) UpdateProfile(ctx context.Context, in *account.UpdateProfileRequest, out *account.UpdateProfileResponse) error {
	return h.profileBiz.UpdateProfile(getTraceContext(ctx, in.GetRequestId(), in.GetProfile().GetUserId()), in, out)
}

func (h *AccountHandlerImpl) Register(ctx context.Context, in *account.RegisterRequest, out *account.RegisterResponse) error {
	return h.accountBiz.Register(getTraceContext(ctx, in.GetRequestId(), 0), in, out)
}

func (h *AccountHandlerImpl) Login(ctx context.Context, in *account.LoginRequest, out *account.LoginResponse) error {
	return h.accountBiz.Login(getTraceContext(ctx, in.GetRequestId(), 0), in, out)
}

func (h *AccountHandlerImpl) Logout(ctx context.Context, in *account.LogoutRequest, out *account.LogoutResponse) error {
	return h.accountBiz.Logout(getTraceContext(ctx, in.GetRequestId(), 0), in, out)
}

func (h *AccountHandlerImpl) Authenticate(ctx context.Context, in *account.AuthRequest, out *account.AuthResponse) error {
	return h.accountBiz.Authenticate(getTraceContext(ctx, in.GetRequestId(), 0), in, out)
}

func getTraceContext(ctx context.Context, requestId string, userId uint64) context.Context {
	return context.WithValue(ctx, "traceKey", map[string]any{
		"request_id": requestId,
		"user_id":    userId,
	})
}
