package db

import (
	at "github.com/lucaseg/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
)

type dbRepository struct {
}

func (db *dbRepository) GetById(id string) (*at.AccessToken, *errors.RestError) {
	return nil, nil
}

func New() *dbRepository {
	return &dbRepository{}
}
