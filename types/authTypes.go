package types

import "time"

type AuthRepository interface {
	IsEmailExist(email string) bool
	IsUsernameExist(email string) bool
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type AuthService interface {
	Signup(params SignupDTO) error
	Login(params LoginDTO) (*LoginResponse, error)
	ForgotPassword(params ForgotPasswordDTO) error
}

type SignupDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email"`
}

type PasswordResetLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"userId" gorm:"not null;index;uniqueIndex:userid_code"`
	Code      int       `json:"code" gorm:"not null;uniqueIndex:userid_code"`
	IsUsed    bool      `json:"isUsed" gorm:"default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;type:TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime;type:TIMESTAMP"`
}
