package rest

import (
	"net/http"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestLoginTimeOut(t *testing.T) {
	rest.StartMockupServer()
	//rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@test.com", "password":"test-password"}`,
		RespHTTPCode: -1,
		RespBody:     "{}",
	})

	userRepo := userRestRepository{}

	user, err := userRepo.Login("email@test.com", "test-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestResponseDistinctOk(t *testing.T) {

}

func TestInvalidJsonError(t *testing.T) {

}

func TestInvalidJsonUser(t *testing.T) {

}

func TestFailedDependency(t *testing.T) {

}

func TestUserLoginSuccess(t *testing.T) {

}
