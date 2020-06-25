package model

import (
	"fmt"
	"time"
)

type TodoContent struct {
	Topic   string     `json:"topic"`
	Detail  *string    `json:"detail"`
	DueDate *time.Time `json:"due_date"`
}

func (t *TodoContent) Validate() error {
	if t.Topic == "" {
		return fmt.Errorf("topic is blank")
	}

	return nil
}

type TodoInfo struct {
	TodoContent
	ID       string    `json:"id"`
	IsDone   bool      `json:"is_done"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
