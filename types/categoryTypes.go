package types

import "time"

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null;size:100;uniqueIndex:name_createdbyid"`
	IsSystem    bool      `json:"isSystem" gorm:"default:0"`
	CreatedById uint      `json:"createdById,omitempty" gorm:"default:0;index;uniqueIndex:name_createdbyid"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime;type:TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime;type:TIMESTAMP"`
}

type CategoryRepository interface {
	IsExist(userId int, name string) bool
	Create(category *Category) error
	GetById(id int) (*Category, error)
	GetAll(userId int, params GetAllCategoryDTO) (*[]Category, error)
	Update(userId int, params UpdateCategoryDTO) error
	Delete(userId int, id int) error
}

type CategoryService interface {
	Create(userId int, params CreateCategoryDTO) error
	GetAll(userId int, params GetAllCategoryDTO) (*[]Category, error)
	Update(userId int, params UpdateCategoryDTO) error
	Delete(userId int, params DeleteCategoryDTO) error
}

type CreateCategoryDTO struct {
	Name     string `json:"name"`
	IsSystem bool   `json:"isSystem"`
}

type GetAllCategoryDTO struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type UpdateCategoryDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DeleteCategoryDTO struct {
	ID int `json:"id"`
}
