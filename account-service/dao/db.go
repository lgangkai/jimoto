package dao

import (
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
