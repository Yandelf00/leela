package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Skill struct {
	Id           uint       `gorm:"primaryKey"`
	Name         string     `gorm:"not null"`
	GoalLevel    string     `gorm:"not null"`
	CurrentLevel int        `gorm:"not null"`
	Practices    []Practice `gorm:"foreignKey:SkillId"`
}

func (s *Skill) BeforeSave(tx *gorm.DB) (err error) {
	if s.GoalLevel != Green && s.GoalLevel != Blue && s.GoalLevel != Orange && s.GoalLevel != Red {
		return fmt.Errorf("invalid role")
	}
	return nil
}
