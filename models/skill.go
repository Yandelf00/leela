package models

import (
	"fmt"

	"gorm.io/gorm"
)

// the Skill model
type Skill struct {
	Id           uint       `gorm:"primaryKey"`
	Name         string     `gorm:"not null"`
	GoalLevel    string     `gorm:"not null"`
	CurrentLevel int        `gorm:"not null"`
	Practices    []Practice `gorm:"foreignKey:SkillId"`
}

// the BeforeSave function ensures that the
// GoalLevel will only be one of the four options
// that we have, namely : green, blue, orange and red
func (s *Skill) BeforeSave(tx *gorm.DB) (err error) {
	if s.GoalLevel != Green && s.GoalLevel != Blue && s.GoalLevel != Orange && s.GoalLevel != Red {
		return fmt.Errorf("invalid role")
	}
	return nil
}
