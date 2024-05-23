package auth

import (
	"fmt"

	"github.com/2marks/go-expense-tracker-api/types"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) IsEmailExist(email string) bool {
	result := r.db.Select("email").Where("email=?", email).First(&types.User{})

	if result.Error != nil {
		return false
	}

	return result.RowsAffected > 0
}

func (r *Repository) IsUsernameExist(username string) bool {
	result := r.db.Select("username").Where("username=?", username).First(&types.User{})

	if result.Error != nil {
		return false
	}

	return result.RowsAffected > 0
}

func (r *Repository) CreateUser(user *types.User) error {
	result := r.db.Create(user)

	if result.Error != nil {
		fmt.Printf("error occured while creating user:%s. err:%s", user.Username, result.Error.Error())
		return fmt.Errorf("error occured while creating user account")
	}

	return nil
}

func (r *Repository) GetUserByUsername(username string) (*types.User, error) {
	user := new(types.User)
	result := r.db.Select("id", "username", "email", "password").Where("username=?", username).First(user)

	if result.Error != nil || result.RowsAffected <= 0 {
		fmt.Printf("error occured while fetching user by username(%s). err: %s", username, result.Error.Error())
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)
	result := r.db.Select("id", "username", "email", "password").Where("email=?", email).First(user)

	if result.Error != nil || result.RowsAffected <= 0 {
		fmt.Printf("error occured while fetching user by email(%s). err: %s", email, result.Error.Error())
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}
