package dao

import (
	"context"
	"gorm.io/gorm"
)

// DBMaster and DBSlave
// Wrapped structs for avoiding wire’s error when there are two same types of input parameters.
type DBMaster struct {
	*gorm.DB
}
type DBSlave struct {
	*gorm.DB
}

type contextTxKey struct{}

func (d *DBMaster) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *DBMaster) db(ctx context.Context) *DBMaster {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return &DBMaster{tx}
	}
	return d
}
