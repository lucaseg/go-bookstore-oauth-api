package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lucaseg/go-bookstore-oauth-api/src/domain/user"
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	restClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}

	userLoginEndpoint = "/users/login"
)

type UserRestRepository interface {
	Login(string, string) (*user.User, *errors.RestError)
}

type userRestRepository struct {
}

func NewUserRestRepository() UserRestRepository {
	return &userRestRepository{}
}

func (s *userRestRepository) Login(email string, password string) (*user.User, *errors.RestError) {
	userLoginRequest := user.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var user user.User
	response := restClient.Post(userLoginEndpoint, userLoginRequest)

	if response == nil || response.Response == nil {
		return nil, errors.InteralServerError("Invalid rest client error trying to login user")
	}

	if response.StatusCode != http.StatusOK {
		var responseError errors.RestError

		if err := json.Unmarshal(response.Bytes(), &responseError); err != nil {
			return nil, errors.InteralServerError("Unexpected response error")
		}

		if response.StatusCode == http.StatusInternalServerError {
			responseError.Status = http.StatusFailedDependency
			return nil, &responseError
		}
	}

	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.InteralServerError("Error trying to unmarshal user login response")
	}

	return &user, nil
}
