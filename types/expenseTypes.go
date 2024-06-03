package types

import "time"

var PaymentMethods []string = []string{"cash", "bank_transfer", "card", "mobile_money", "paypal"}

type Expense struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"userId" gorm:"index"`
	CategoryId    uint      `json:"categoryId" gorm:"not null"`
	Category      string    `json:"category" gorm:"not null;size:100"`
	PaymentMethod string    `json:"paymentMethod" gorm:"not null;size:100"`
	Amount        float64   `json:"amount" gorm:"not null"`
	Currency      string    `json:"currency" gorm:"size:5;not null"`
	Description   string    `json:"description" gorm:"not null"`
	PayedAt       time.Time `json:"payedAt" gorm:"type:TIMESTAMP"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime;type:TIMESTAMP"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime;type:TIMESTAMP"`
}

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null;size:100;uniqueIndex:name_createdbyid"`
	IsSystem    bool   `gorm:"default:1"`
	CreatedById uint   `gorm:"default:0;index;uniqueIndex:name_createdbyid"`
}

type CreateExpenseDTO struct {
	Category      string    `json:"category"`
	PaymentMethod string    `json:"paymentMethod"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Description   string    `json:"description"`
	PayedAt       time.Time `json:"payedAt"`
}

type UpdateExpenseDTO struct {
	ID int `json:"id"`
	CreateExpenseDTO
}

type DeleteExpenseDTO struct {
	ID int `json:"id"`
}

type GetAllExpensesDTO struct {
	Page        int `json:"page"`
	PerPage     int `json:"perPage"`
	SearchQuery int `json:"searchQuery"`
}

type GetOneExpenseDTO struct {
	ID int `json:"id"`
}

type ExpenseRepository interface {
	Create(expense *Expense) error
	GetCategoryByName(name string) (*Category, error)
	GetAll(userId int, params GetAllExpensesDTO) (*[]Expense, error)
	GetById(userId int, id int) (*Expense, error)
	Update(userId int, params UpdateExpenseDTO) error
	Delete(userId int, id int) error
	GetUserDefaultCurrency(userId int) string
}

type ExpenseService interface {
	Create(userId int, params CreateExpenseDTO) error
	GetAll(userId int, params GetAllExpensesDTO) (*[]Expense, error)
	GetById(userId int, params GetOneExpenseDTO) (*Expense, error)
	Update(userId int, params UpdateExpenseDTO) error
	Delete(userId int, params DeleteExpenseDTO) error
}
