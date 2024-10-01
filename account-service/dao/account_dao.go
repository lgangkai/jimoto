package dao

import (
	"context"
	"github.com/lgangkai/golog"
	"jimoto/account-service/model"
)

type AccountDao struct {
	db     *DBMaster
	logger *golog.Logger
}

func NewAccountDao(db *DBMaster, logger *golog.Logger) *AccountDao {
	return &AccountDao{
		db:     db,
		logger: logger,
	}
}

func (d *AccountDao) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	d.logger.Info(ctx, "Call AccountDao.GetUserByEmail, email: ", email)
	user := &model.User{}
	if err := d.db.Where("email = ?", email).Take(user).Error; err != nil {
		d.logger.Error(ctx, "Fail to get data, err: ", err.Error())
		return nil, err
	}
	d.logger.Info(ctx, "Get user done, user: ", *user)
	return user, nil
}

func (d *AccountDao) Insert(ctx context.Context, user *model.User) error {
	d.logger.Info(ctx, "Call AccountDao.Insert, user: ", user)
	if err := d.db.Create(user).Error; err != nil {
		d.logger.Error(ctx, "Fail to insert into sql DB, err: ", err.Error())
		return err
	}
	d.logger.Info(ctx, "Insert data into sql DB succeed.")
	return nil
}
