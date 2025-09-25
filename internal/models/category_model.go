package model

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null:unique"`
	Description string    `gorm:"default:null"`
	CreatedAt   time.Time `gorm:"type:timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
}
