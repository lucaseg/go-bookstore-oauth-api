package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucaseg/go-bookstore-oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	c.JSON(http.StatusFailedDependency, "implement me")
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
