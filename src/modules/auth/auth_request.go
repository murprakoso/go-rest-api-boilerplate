package auth

type SRegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SLoginRequest struct {
	Username string `json:"username" binding:"required"`
	//Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
