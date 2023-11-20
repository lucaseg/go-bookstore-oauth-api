package access_token

import (
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(key string) (*AccessToken, *errors.RestError) {
	if key == "" {
		return nil, errors.BadRequest("The key can not be empty")
	}

	accessToken, err := s.repository.GetById(key)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
