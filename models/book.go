package models

// the Book model
type Book struct {
	Id       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Category string
	Author   string
	IsRead   bool `gorm:"not null"`
	Rating   int
}
