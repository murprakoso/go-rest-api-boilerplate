package auth

import (
	"time"
)

type SAuthResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SAuthDetailResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAuthDetailResponseFromEntity(auth User) SAuthDetailResponse {
	return SAuthDetailResponse{
		ID:        auth.ID,
		Name:      auth.Name,
		Username:  auth.Username,
		Email:     auth.Email,
		Password:  auth.Password,
		CreatedAt: auth.CreatedAt,
		UpdatedAt: auth.UpdatedAt,
	}
}
