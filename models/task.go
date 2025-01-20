package models

import "time"

// the Task model
type Task struct {
	Id         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Type       string `gorm:"not null"`
	Difficulty string `gorm:"not null"`
	IsDone     bool   `gorm:"not null"`
	CreatedAt  time.Time
}
