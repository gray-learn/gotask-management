package models

import (
	"time"
)

type Task struct {
	ID          int64
	Title       string
	Description string
	Status      string
	AssignedTo  int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deadline    time.Time
}