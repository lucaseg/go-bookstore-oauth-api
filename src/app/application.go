package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaseg/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/lucaseg/go-bookstore-oauth-api/src/http"
	"github.com/lucaseg/go-bookstore-oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atRepository := db.New()
	service := access_token.NewService(atRepository)
	accessTokenHandler := http.NewAccessTokenHandler(service)

	router.GET("", accessTokenHandler.GetById)
	router.Run()
}
