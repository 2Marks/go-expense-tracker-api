package users

import (
	"fmt"
	"strings"

	"github.com/2marks/go-expense-tracker-api/types"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetById(id int) (*types.User, error) {
	user := new(types.User)
	result := r.db.Omit("password").First(user, id)

	if result.Error != nil {
		fmt.Printf("error occured while fetching user by id(%v). err: %s", id, result.Error.Error())
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *Repository) UpdateUserDetails(params types.UpdateUserDetailsDTO) error {
	result := r.db.Model(&types.User{}).Where("id=?", params.ID).Updates(types.User{
		Name:            params.Name,
		DefaultCurrency: strings.ToUpper(params.DefaultCurrency),
	})

	if result.Error != nil {
		fmt.Printf("error while updating user with id :%v. err:%s", params.ID, result.Error.Error())
		return fmt.Errorf("error while updating user with id :%v", params.ID)
	}

	return nil
}
