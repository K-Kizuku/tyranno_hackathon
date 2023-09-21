package model

import "time"

type Post struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"not null" json:"content"`
	User      User      `gorm:"not null;foreignkey:Author"`
	Author    int       `gorm:"not null" json:"author"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
