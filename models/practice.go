package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Practice struct {
	Id         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Difficulty string `gorm:"not null"`
	SkillId    uint   `gorm:"not null"`
}

func (s *Practice) BeforeSave(tx *gorm.DB) (err error) {
	if s.Difficulty != Easy && s.Difficulty != Moderate && s.Difficulty != Hard {
		return fmt.Errorf("invalid difficulty")
	}
	return nil
}
