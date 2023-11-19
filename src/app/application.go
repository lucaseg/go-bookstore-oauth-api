package app

import (
	"github.com/lucaseg/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/lucaseg/go-bookstore-oauth-api/src/repository/db"
)

func StartApplication() {
	atRepository := db.New()
	service := access_token.NewService(atRepository)
}
