package db

import (
	"github.com/gocql/gocql"
	"github.com/lucaseg/go-bookstore-oauth-api/src/clients/cassandra"
	at "github.com/lucaseg/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM oauth.access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO oauth.access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateAccessToken = "UPDATE oauth.access_tokens SET expires=? WHERE access_token=?"
)

type dbRepository struct {
}

func New() *dbRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetById(id string) (*at.AccessToken, *errors.RestError) {
	// get session cassandra db
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var result at.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.Token,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NotFound("The access token not found")
		}
		return nil, errors.InteralServerError(err.Error())
	}

	return &result, nil
}

func (db *dbRepository) Create(at *at.AccessToken) *errors.RestError {
	// get session cassandra db
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		at.Token,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.InteralServerError(err.Error())
	}
	return nil
}

func (db *dbRepository) UpdateExpirationTime(at *at.AccessToken) *errors.RestError {
	// get session cassandra db
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	if err := session.Query(queryUpdateAccessToken,
		at.Token,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.InteralServerError(err.Error())
	}
	return nil
}
