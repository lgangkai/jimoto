package dao

import (
	"context"
	"github.com/lgangkai/golog"
	"jimoto/account-service/model"
)

type UserDao struct {
	db     *DBMaster
	logger *golog.Logger
}

func NewUserDao(db *DBMaster, logger *golog.Logger) *UserDao {
	return &UserDao{
		db:     db,
		logger: logger,
	}
}

func (d *UserDao) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	d.logger.Info(ctx, "Call UserDao.GetUserByEmail, email: ", email)
	user := &model.User{}
	if err := d.db.Where("email = ?", email).Take(user).Error; err != nil {
		d.logger.Error(ctx, "Fail to get data, err: ", err.Error())
		return nil, err
	}
	d.logger.Info(ctx, "Get user done, user: ", *user)
	return user, nil
}

func (d *UserDao) Insert(ctx context.Context, user *model.User) error {
	d.logger.Info(ctx, "Call UserDao.Insert, user: ", user)
	if err := d.db.db(ctx).Create(user).Error; err != nil {
		d.logger.Error(ctx, "Fail to insert into sql DB, err: ", err.Error())
		return err
	}
	d.logger.Info(ctx, "Insert data into sql DB succeed.")
	return nil
}
