package model

import "time"

type Task struct {
	ID        int64
	UserId    int64
	Status    int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
