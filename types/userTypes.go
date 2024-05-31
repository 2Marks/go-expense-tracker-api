package types

import "time"

type User struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"size:100;not null"`
	Username        string    `json:"username" gorm:"unique;size:100;not null"`
	Email           string    `json:"email" gorm:"unique;not null"`
	Password        string    `json:"password,omitempty"`
	IsActive        bool      `json:"isActive" gorm:"default:1"`
	DefaultCurrency string    `json:"defaultCurrency" gorm:"size:5"`
	CreatedAt       time.Time `json:"createdAt" gorm:"autoCreateTime;type:TIMESTAMP"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"autoUpdateTime;type:TIMESTAMP"`
}

type UserRepository interface {
	GetById(id int) (*User, error)
	UpdateUserDetails(params UpdateUserDetailsDTO) error
}

type UserService interface {
	GetLoggedInUser(id int) (*User, error)
	UpdateUserDetails(params UpdateUserDetailsDTO) error
}

type UpdateUserDetailsDTO struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	DefaultCurrency string `json:"defaultCurrency"`
}
