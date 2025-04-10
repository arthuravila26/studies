package service

import (
	"github.com/arthuravila26/studies/go-project-1/go-gateway-api/internal/domain"
	"github.com/arthuravila26/studies/go-project-1/go-gateway-api/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) *AccountService {
	return &AccountService{repository: repo}
}


func (s *AccountService) CreateAccount(input dto.CreateAccount) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}
	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return output, nil
}

func (s *AccountService) FindByID(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return output, nil
}
