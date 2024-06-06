package expenses

import (
	"fmt"
	"strings"
	"time"

	"github.com/2marks/go-expense-tracker-api/errs"
	"github.com/2marks/go-expense-tracker-api/types"
)

type Service struct {
	repo types.ExpenseRepository
}

func NewService(repo types.ExpenseRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(userId int, params types.CreateExpenseDTO) error {
	category, err := s.repo.GetCategoryByName(params.Category)
	if err != nil {
		return errs.ErrResourceNotFound(err)
	}

	ok, paymentMethodStr := isValidPaymentMethod(params.PaymentMethod)
	if !ok {
		return errs.ErrUnprocessableEntity(fmt.Errorf("supplied payment method(%s) not valid. valid options includes: %s", params.PaymentMethod, paymentMethodStr))
	}

	amountCurrency := params.Currency
	if amountCurrency == "" {
		amountCurrency = s.repo.GetUserDefaultCurrency(userId)
	}

	if amountCurrency == "" {
		return errs.ErrUnprocessableEntity(fmt.Errorf("kindly supply the amount currency to continue"))
	}

	expense := &types.Expense{
		UserID:        uint(userId),
		CategoryId:    category.ID,
		Category:      category.Name,
		PaymentMethod: params.PaymentMethod,
		Amount:        params.Amount,
		Currency:      strings.ToUpper(amountCurrency),
		Description:   params.Description,
		PayedAt:       time.Now(), //params.PayedAt,
	}

	return s.repo.Create(expense)
}

func (s *Service) GetAll(userId int, params types.GetAllExpensesDTO) (*[]types.Expense, error) {
	return s.repo.GetAll(userId, params)
}

func (s *Service) GetById(userId int, params types.GetOneExpenseDTO) (*types.Expense, error) {
	return s.repo.GetById(userId, params.ID)
}

func (s *Service) Update(userId int, params types.UpdateExpenseDTO) error {
	_, err := s.repo.GetById(userId, params.ID)

	if err != nil {
		return err
	}

	return s.repo.Update(userId, params)
}

func (s *Service) Delete(userId int, params types.DeleteExpenseDTO) error {
	_, err := s.repo.GetById(userId, params.ID)

	if err != nil {
		return err
	}

	return s.repo.Delete(userId, params.ID)
}

func isValidPaymentMethod(paymentMethod string) (bool, string) {
	isValid := false

	for _, method := range types.PaymentMethods {
		if paymentMethod == method {
			isValid = true
			break
		}
	}

	return isValid, strings.Join(types.PaymentMethods, ",")
}
