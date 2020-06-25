package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"todo/pkg/util"
)

type Base struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (Base) BeforeCreate(scope *gorm.Scope) error {
	id, ok := scope.FieldByName("ID")
	if !ok {
		return fmt.Errorf("not found column id")
	}
	if s := id.Field.String(); !ok || s == "" {
		return scope.SetColumn("ID", util.NewID())
	}
	return nil
}
