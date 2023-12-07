package access_token

import (
	"github.com/lucaseg/go-bookstore-oauth-api/src/repository/rest"
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(*AccessToken) *errors.RestError
	UpdateExpirationTime(*AccessToken) *errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(*AccessTokenRequest) (*AccessToken, *errors.RestError)
	UpdateExpirationTime(*AccessToken) *errors.RestError
}

type service struct {
	repository     Repository
	userRepository rest.UserRestRepository
}

func NewService(repo Repository, userRepository rest.UserRestRepository) Service {
	return &service{
		repository:     repo,
		userRepository: userRepository,
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

func (s *service) Create(at *AccessTokenRequest) (*AccessToken, *errors.RestError) {
	if err := at.Validate(); err != nil {
		return nil, err
	}

	// Login user using users service api
	user, err := s.userRepository.Login(at.Email, at.Password)
	if err != nil {
		return nil, err
	}

	// if login was success we will create the access token
	newAccessToken := GetNewAccessToken(at.GrantType, user.Id)
	newAccessToken.Generate()

	if err := s.repository.Create(newAccessToken); err != nil {
		return nil, err
	}

	return newAccessToken, nil
}

func (s *service) UpdateExpirationTime(at *AccessToken) *errors.RestError {
	return nil
}
