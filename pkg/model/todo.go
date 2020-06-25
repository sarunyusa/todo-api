package model

import "time"

type TodoContent struct {
	Topic   string    `json:"topic"`
	Detail  string    `json:"detail"`
	DueDate time.Time `json:"due_date"`
}

type TodoInfo struct {
	TodoContent
	ID       string    `json:"id"`
	IsDone   bool      `json:"is_done"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
