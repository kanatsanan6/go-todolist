package model

import "time"

type Task struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}
