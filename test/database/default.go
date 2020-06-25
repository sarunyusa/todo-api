package database

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func NewSqlMockConnection() (*gorm.DB, sqlmock.Sqlmock) {
	db, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("sqlite3", db)
	gormDB.LogMode(true)
	return gormDB, sqlMock
}
