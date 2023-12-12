package access_token

import (
	"fmt"
	"time"

	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/crypto_utils"
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
)

const (
	expirationTime       = 24
	grantTypePassword    = "password"
	grantTypeCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Expires      int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	UserId       int64  `json:"userId"`
	ClientId     string `json:"client_id"`
}

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestError {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeCredentials:
		break
	default:
		return errors.InteralServerError("Invalid access token type")
	}
	return nil
}

func GetNewAccessToken(grantType string, userId int64) *AccessToken {
	return &AccessToken{
		TokenType:    grantType,
		UserId:       userId,
		Expires:      time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
		RefreshToken: "",
	}
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestError {
	if at.AccessToken == "" {
		return errors.BadRequest("Invalid access token value")
	}

	return nil
}
