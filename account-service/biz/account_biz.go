package biz

import (
	"context"
	"github.com/lgangkai/golog"
	"jimoto/account-service/service"
	"protos/account"
)

type AccountBiz struct {
	accountService *service.AccountService
	logger         *golog.Logger
}

func NewAccountBiz(accountService *service.AccountService, logger *golog.Logger) *AccountBiz {
	return &AccountBiz{
		accountService: accountService,
		logger:         logger,
	}
}

func (b *AccountBiz) Register(ctx context.Context, in *account.RegisterRequest, out *account.RegisterResponse) error {
	b.logger.Info(ctx, "Call AccountBiz.Register, request: ", in)
	err := b.accountService.Register(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		b.logger.Error(ctx, "Register failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call AccountBiz.Register successfully.")
	return nil
}

func (b *AccountBiz) Login(ctx context.Context, in *account.LoginRequest, out *account.LoginResponse) error {
	b.logger.Info(ctx, "Call AccountBiz.Login, request: ", in)
	token, userId, err := b.accountService.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		b.logger.Error(ctx, "Login failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call AccountBiz.Login successfully.")
	out.Token = *token
	out.UserId = *userId
	return nil
}

// Logout Currently no implement for logout in service.
// Just remove token in api gateway.
func (b *AccountBiz) Logout(ctx context.Context, in *account.LogoutRequest, out *account.LogoutResponse) error {
	return nil
}

func (b *AccountBiz) Authenticate(ctx context.Context, in *account.AuthRequest, out *account.AuthResponse) error {
	b.logger.Info(ctx, "Call AccountBiz.Authenticate, request: ", in)
	userId, email, err := b.accountService.Authenticate(ctx, in.GetToken())
	if err != nil {
		b.logger.Error(ctx, "Authenticate failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call AccountBiz.Authenticate successfully.")
	out.UserId = userId
	out.Email = email
	return nil
}
