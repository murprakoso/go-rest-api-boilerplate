package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successfully",
		"success":     true,
		"data":        user,
		"accessToken": token,
	})
}

func (h *SAuthHandler) Validate(c *gin.Context) {
	//claims, _ := c.Get("claims")
	claims := c.MustGet("claims").(jwt.MapClaims)

	// Dalam JWT payload, username mungkin berada di bawah key tertentu, sesuaikan dengan struktur payload Anda
	username, ok := claims["username"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Username tidak ditemukan dalam claims",
		})
		return
	}

	user, _ := h.authService.FindByUsername(username)

	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
		"data":    NewAuthDetailResponseFromEntity(user),
	})
}
