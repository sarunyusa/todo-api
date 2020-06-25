package entity

import "time"

type Todo struct {
	Base
	Topic   string `gorm:"not null;"`
	Detail  *string
	DueDate *time.Time
	IsDone  bool `gorm:"not null;"`
}
