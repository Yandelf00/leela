package models

import (
	"fmt"

	"gorm.io/gorm"
)

// the Practice model
type Practice struct {
	Id         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Difficulty string `gorm:"not null"`
	SkillId    uint   `gorm:"not null"`
}

// the BeforeSave function ensures that the
// difficulty will be one of the 3 options that
// chose, namely : easy, moderate and hard
func (s *Practice) BeforeSave(tx *gorm.DB) (err error) {
	if s.Difficulty != Easy && s.Difficulty != Moderate && s.Difficulty != Hard {
		return fmt.Errorf("invalid difficulty")
	}
	return nil
}
