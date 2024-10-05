package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	errs "errs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lgangkai/golog"
	"gorm.io/gorm"
	"jimoto/account-service/dao"
	"jimoto/account-service/model"
	"time"
)

const (
	LOGIN_EXPIRE_TIME_HOUR = 24 * time.Hour * 30
	JWT_SIGNING_KEY        = "ofB1pXMJFKs9N11yXomfPr1Vq0h5GE80"
)

type UserClaim struct {
	jwt.RegisteredClaims
	UserId uint64 `json:"user_id"`
	Email  string `json:"email"`
}

type AccountService struct {
	accountDao *dao.UserDao
	profileDao *dao.ProfileDao
	db         *dao.DBMaster
	logger     *golog.Logger
}

func NewAccountService(accountDao *dao.UserDao, profileDao *dao.ProfileDao, db *dao.DBMaster, logger *golog.Logger) *AccountService {
	return &AccountService{
		accountDao: accountDao,
		profileDao: profileDao,
		db:         db,
		logger:     logger,
	}
}

func (s *AccountService) Register(ctx context.Context, email string, password string) error {
	s.logger.Info(ctx, "Call AccountService.Register, email: ", email)
	// 1. check whether email has been registered.
	user, err := s.accountDao.GetUserByEmail(ctx, email)
	//   1.1 if user exists, return error.
	if user != nil && user.Id != 0 {
		s.logger.Error(ctx, "This email has already been registered.")
		return errs.New(errs.ERR_EMAIL_IS_REGISTERED)
	}
	//   1.2 if err is sql DB internal error, return error.
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error(ctx, "sql DB internal error, err: ", err.Error())
		return errs.New(errs.ERR_REGISTER_INTERNAL)
	}

	// 2. save email and password. and create a default profile
	user = &model.User{
		Password: pswMd5Encode(password),
		Email:    email,
	}
	// start transaction
	return s.db.ExecTx(ctx, func(ctx context.Context) error {
		err = s.accountDao.Insert(ctx, user)
		// get id after insert
		err = s.profileDao.Insert(ctx, &model.Profile{
			UserId: user.Id,
			Email:  email,
		})
		if err != nil {
			s.logger.Error(ctx, "Insert user/profile failed, err: ", err.Error())
			return errs.New(errs.ERR_REGISTER_INTERNAL)
		}
		return nil
	})
}

func (s *AccountService) Login(ctx context.Context, email string, password string) (*string, *uint64, error) {
	s.logger.Info(ctx, "Call AccountService.Login, email: ", email)
	// 1. verify email and password.
	user, err := s.accountDao.GetUserByEmail(ctx, email)
	if err != nil {
		s.logger.Error(ctx, "Get user failed, err: ", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errs.New(errs.ERR_LOGIN_NO_USER)
		} else {
			return nil, nil, errs.New(errs.ERR_LOGIN_INTERNAL)
		}
	}
	if user.Password != pswMd5Encode(password) {
		s.logger.Error(ctx, "Password mismatch.")
		return nil, nil, errs.New(errs.ERR_PASSWORD_MISMATCH)
	}
	// 2. login succeed, return token.
	claim := &UserClaim{}
	claim.UserId = user.Id
	claim.Email = email
	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRE_TIME_HOUR))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(JWT_SIGNING_KEY))
	if err != nil {
		s.logger.Error(ctx, "Generate token failed, err: ", err.Error())
		return nil, nil, errs.New(errs.ERR_LOGIN_INTERNAL)
	}
	s.logger.Info(ctx, "Call AccountService.Login succeed.")
	return &tokenString, &user.Id, nil
}

func (s *AccountService) Authenticate(ctx context.Context, token string) (uint64, string, error) {
	s.logger.Info(ctx, "Call AccountService.Authenticate, token: ", token)
	claim := &UserClaim{}
	tk, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SIGNING_KEY), nil
	})
	if err != nil {
		s.logger.Error(ctx, "Parse token failed, err: ", err.Error())
		return 0, "", errs.New(errs.ERR_AUTH_FAILED)
	}
	if tk == nil || !tk.Valid || claim.UserId == 0 {
		s.logger.Error(ctx, "Token invalid.")
		return 0, "", errs.New(errs.ERR_AUTH_FAILED)
	}

	// check whether expired.
	if time.Now().After(claim.ExpiresAt.Time) {
		s.logger.Error(ctx, "Token expired.")
		return 0, "", errs.New(errs.ERR_TOKEN_EXPIRED)
	}
	s.logger.Info(ctx, "Call AccountService.Authenticate succeed.")
	return claim.UserId, claim.Email, nil
}

func pswMd5Encode(psw string) string {
	data := []byte(psw)
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}
