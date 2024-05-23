package types

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Username  string    `json:"username" gorm:"unique;size:100;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"isActive" gorm:"default:1"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;type:TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime;type:TIMESTAMP"`
}
