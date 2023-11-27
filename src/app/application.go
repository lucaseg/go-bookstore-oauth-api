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
	//cassandraClient := cassandra.NewClient()
	service := access_token.NewService(atRepository)
	accessTokenHandler := http.NewAccessTokenHandler(service)

	router.GET("/oauth/access-token/:access-token-id", accessTokenHandler.GetById)
	router.POST("oauth/access-token", accessTokenHandler.Create)
	router.POST("oauth/access-token/:access-token-id", accessTokenHandler.Update)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
