package access_token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstantsAccessToken(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.True(t, at.Client == 0, "The new access token should have client 0")
	assert.True(t, at.UserId == 0, "The user id should be 0")
	assert.True(t, at.AccessToken == "", "Should be empty access token")
}

func TestExpirationTime(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "The access token should not be expired")
}
