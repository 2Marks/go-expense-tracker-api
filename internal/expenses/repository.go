package expenses

import (
	"fmt"

	"github.com/2marks/go-expense-tracker-api/errors"
	"github.com/2marks/go-expense-tracker-api/types"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(expense *types.Expense) error {
	result := r.db.Create(expense)

	if result.Error != nil {
		fmt.Printf("error while creating expense. err:%s \n", result.Error.Error())
		return fmt.Errorf("error while creating expense")
	}

	return nil
}

func (r *Repository) GetCategoryByName(name string) (*types.Category, error) {
	category := new(types.Category)
	result := r.db.Select("id", "name").Where("name=? and is_system=?", name, 1).Or("name=? and is_system=?", name, 0).First(category)

	if result.Error != nil || result.RowsAffected <= 0 {
		return nil, fmt.Errorf("category not found")
	}

	return category, nil
}

func (r *Repository) GetAll(userId int, params types.GetAllExpensesDTO) (*[]types.Expense, error) {
	limit := params.PerPage
	offset := (params.Page - 1) * params.PerPage

	expenses := make([]types.Expense, 0)
	result := r.db.Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&expenses)

	if result.Error != nil {
		fmt.Printf("error while fetching expenses for user:%d. err:%s", userId, result.Error.Error())
		return nil, fmt.Errorf("error while fetching expenses")
	}

	return &expenses, nil
}

func (r *Repository) GetById(userId int, id int) (*types.Expense, error) {
	expense := new(types.Expense)
	result := r.db.Where("user_id=? AND id=?", userId, id).First(expense)

	if result.Error != nil {
		fmt.Printf("error while fetching expense by id. err:%s \n", result.Error.Error())
		return nil, errors.ErrResourceNotFound(fmt.Errorf("expense not found"))
	}

	return expense, nil
}

func (r *Repository) Update(userId int, params types.UpdateExpenseDTO) error {
	expenseToUpdate := types.Expense{
		PaymentMethod: params.PaymentMethod,
		Amount:        params.Amount,
		Description:   params.Description,
		PayedAt:       params.PayedAt,
	}

	result := r.db.Model(&types.Expense{}).Where("user_id=? AND id=?", userId, params.ID).Omit("UserID", "CategoryId", "Category").Updates(expenseToUpdate)

	if result.Error != nil {
		fmt.Printf("error while updating expense with id:%v. err:%s \n", params.ID, result.Error.Error())
		return fmt.Errorf("error while updating expense")
	}

	return nil
}

func (r *Repository) Delete(userId int, id int) error {
	result := r.db.Where("user_id=? AND id=?", userId, id).Delete(&types.Expense{})

	if result.Error != nil {
		fmt.Printf("error while deleting expense with id:%v. err:%s \n", id, result.Error.Error())
		return fmt.Errorf("error while deleting expense")
	}

	return nil
}

func (r *Repository) GetUserDefaultCurrency(userId int) string {
	user := new(types.User)
	result := r.db.Select("default_currency").First(user, userId)

	if result.Error != nil {
		fmt.Printf("error while fetching user:%v default currency. err:%s \n", userId, result.Error.Error())
		return ""
	}

	return user.DefaultCurrency
}
