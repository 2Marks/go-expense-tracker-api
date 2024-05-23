package types

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
