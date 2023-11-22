package access_token

import (
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(*AccessToken) *errors.RestError
	UpdateExpirationTime(*AccessToken) *errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(*AccessToken) *errors.RestError
	UpdateExpirationTime(*AccessToken) *errors.RestError
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

func (s *service) Create(at *AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}

	newAccessToken := GetNewAccessToken()
	newAccessToken.Token = at.Token
	newAccessToken.UserId = at.UserId
	newAccessToken.ClientId = at.ClientId

	if err := s.repository.Create(at); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateExpirationTime(at *AccessToken) *errors.RestError {
	return nil
}
