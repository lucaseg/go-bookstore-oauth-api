package access_token

import (
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	Token    string `json:"token"`
	UserId   int64  `json:"user_id"`
	ClientId int64  `json:"client_id"`
	Expires  int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestError {
	if at.Token == "" {
		return errors.BadRequest("Invalid access token value")
	}

	if at.UserId == 0 {
		return errors.BadRequest("Invalid user_id value")
	}

	if at.ClientId == 0 {
		return errors.BadRequest("Invalid client value")
	}

	return nil
}