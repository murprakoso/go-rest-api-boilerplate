package auth

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-boilerplate/src/commons/core"
	"net/http"
)

type SAuthHandler struct {
	authService IAuthService
}

func NewAuthHandler(authService IAuthService) *SAuthHandler {
	return &SAuthHandler{authService}
}

func (h *SAuthHandler) Register(c *gin.Context) {
	var authRequest SRegisterRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		core.ResponseError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.authService.Register(authRequest)
	if err != nil {
		core.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	core.ResponseJSON(c, http.StatusCreated, true, "", NewAuthDetailResponseFromEntity(user))
}

func (h *SAuthHandler) Login(c *gin.Context) {
	var authRequest SLoginRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		core.ResponseError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, token, err := h.authService.Login(authRequest)
	if err != nil {
		core.ResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successfully",
		"success":     true,
		"data":        user,
		"accessToken": token,
	})
}
