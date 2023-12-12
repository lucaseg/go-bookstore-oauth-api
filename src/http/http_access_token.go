package http

import (
	"net/http"

	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/lucaseg/go-bookstore-oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Update(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	key := c.Param("access-token-id")
	accessToken, err := h.service.GetById(key)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var accessToken access_token.AccessTokenRequest
	if err := c.BindJSON(&accessToken); err != nil {
		restErr := errors.BadRequest("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}

	at, err := h.service.Create(&accessToken)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}

func (h *accessTokenHandler) Update(c *gin.Context) {
	var accessToken access_token.AccessToken
	if err := c.BindJSON(&accessToken); err != nil {
		restErr := errors.BadRequest("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}

	h.service.UpdateExpirationTime(&accessToken)
	c.JSON(http.StatusOK, "")
}
