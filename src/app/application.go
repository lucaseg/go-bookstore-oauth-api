package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaseg/go-bookstore-oauth-api/src/clients/cassandra"
	"github.com/lucaseg/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/lucaseg/go-bookstore-oauth-api/src/http"
)

var (
	router = gin.Default()
)

func StartApplication() {
	//atRepository := db.New()
	cassandraClient := cassandra.NewClient()
	service := access_token.NewService(cassandraClient)
	accessTokenHandler := http.NewAccessTokenHandler(service)

	router.GET("/oauth/access-token/:access-token-id", accessTokenHandler.GetById)
	router.Run()
}
