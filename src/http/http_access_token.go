package http

import (
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
	"net/http"

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
	key := c.Param("key")
	accessToken, err := h.service.GetById(key)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var accessToken access_token.AccessToken
	if err := c.BindJSON(&accessToken); err != nil {
		restErr := errors.BadRequest("invalid json")
		c.JSON(restErr.Status, restErr)
	}

	h.service.Create(&accessToken)
	c.JSON(http.StatusCreated, "{}")
}

func (h *accessTokenHandler) Update(c *gin.Context) {
	var accessToken access_token.AccessToken
	if err := c.BindJSON(&accessToken); err != nil {
		restErr := errors.BadRequest("invalid json")
		c.JSON(restErr.Status, restErr)
	}

	h.service.UpdateExpirationTime(&accessToken)
	c.JSON(http.StatusOK, "")
}