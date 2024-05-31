package users

import "github.com/2marks/go-expense-tracker-api/types"

type Service struct {
	repo types.UserRepository
}

func NewService(repo types.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetLoggedInUser(id int) (*types.User, error) {
	user, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateUserDetails(params types.UpdateUserDetailsDTO) error {
	_, err := s.repo.GetById(params.ID)
	if err != nil {
		return err
	}

	return s.repo.UpdateUserDetails(params)
}
