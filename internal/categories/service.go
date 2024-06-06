package categories

import (
	"fmt"

	"github.com/2marks/go-expense-tracker-api/errs"
	"github.com/2marks/go-expense-tracker-api/types"
)

type Service struct {
	repo types.CategoryRepository
}

func NewService(repo types.CategoryRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(userId int, params types.CreateCategoryDTO) error {
	if isExist := s.repo.IsExist(userId, params.Name); isExist {
		return errs.ErrUnprocessableEntity(fmt.Errorf("category %s already exist", params.Name))
	}

	category := &types.Category{
		Name:        params.Name,
		IsSystem:    false,
		CreatedById: uint(userId),
	}

	return s.repo.Create(category)
}

func (s *Service) GetAll(userId int, params types.GetAllCategoryDTO) (*[]types.Category, error) {
	return s.repo.GetAll(userId, params)
}

func (s *Service) Update(userId int, params types.UpdateCategoryDTO) error {
	category, err := s.repo.GetById(params.ID)
	if err != nil {
		return err
	}

	if category.IsSystem {
		return errs.ErrAccessDenied(fmt.Errorf("you are not allowed to update system defined categories"))
	}

	if category.CreatedById != uint(userId) {
		return errs.ErrResourceNotFound(fmt.Errorf("category not found"))
	}

	return s.Update(userId, params)
}

func (s *Service) Delete(userId int, params types.DeleteCategoryDTO) error {
	category, err := s.repo.GetById(params.ID)
	if err != nil {
		return err
	}

	if category.IsSystem {
		return errs.ErrAccessDenied(fmt.Errorf("you are not allowed to delete system defined categories"))
	}

	if category.CreatedById != uint(userId) {
		return errs.ErrResourceNotFound(fmt.Errorf("category not found"))
	}

	return s.repo.Delete(userId, params.ID)
}
