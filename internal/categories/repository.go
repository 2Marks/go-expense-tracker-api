package categories

import (
	"fmt"

	"github.com/2marks/go-expense-tracker-api/errs"
	"github.com/2marks/go-expense-tracker-api/types"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) IsExist(userId int, name string) bool {
	var count int64
	r.db.Model(&types.Category{}).Where("created_by_id=? AND name=?", userId, name).Or("name=? AND is_system=1", name).Count(&count)

	return count > 0

}

func (r *Repository) Create(category *types.Category) error {
	fmt.Println(*category)
	result := r.db.Create(category)

	if result.Error != nil {
		fmt.Printf("error while creating category. err:%s \n", result.Error.Error())
		return errs.ErrInternalServerError(fmt.Errorf("error while creating category"))
	}

	return nil
}

func (r *Repository) GetAll(userId int, params types.GetAllCategoryDTO) (*[]types.Category, error) {
	limit := params.PerPage
	offset := (params.Page - 1) * params.PerPage

	categories := make([]types.Category, 0)
	result := r.db.Where("created_by_id=? AND is_system=?", userId, 0).Or("is_system=1").Limit(limit).Offset(offset).Omit("created_by_id").Find(&categories)

	if result.Error != nil {
		fmt.Printf("error while fetching categories for user:%d. err:%s", userId, result.Error.Error())
		return nil, errs.ErrInternalServerError(fmt.Errorf("error while fetching categories"))
	}

	return &categories, nil
}

func (r *Repository) GetById(id int) (*types.Category, error) {
	category := new(types.Category)
	result := r.db.First(category, id)

	if result.RowsAffected <= 0 {
		return nil, errs.ErrResourceNotFound(fmt.Errorf("category not found"))
	}

	if result.Error != nil {
		fmt.Printf("error while fetching category by id:%v. err:%s", id, result.Error.Error())
		return nil, errs.ErrInternalServerError(fmt.Errorf("error while fetching category with id:%v", id))
	}

	return category, nil
}

func (r *Repository) Update(userId int, params types.UpdateCategoryDTO) error {
	result := r.db.Model(&types.Category{}).Where("user_id=? AND id=?", userId, params.ID).Update("name", params.Name)

	if result.Error != nil {
		fmt.Printf("error while updating category. err:%s \n", result.Error.Error())
		return errs.ErrInternalServerError(fmt.Errorf("error while updating category"))
	}

	return nil
}

func (r *Repository) Delete(userId int, id int) error {
	result := r.db.Where("created_by_id=? AND id=?", userId, id).Delete(&types.Category{})

	if result.Error != nil {
		fmt.Printf("error while deleting category. err:%s \n", result.Error.Error())
		return errs.ErrInternalServerError(fmt.Errorf("error while deleting category"))
	}

	return nil
}
