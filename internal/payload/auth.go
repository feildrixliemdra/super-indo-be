package payload

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"johndoe@gmail.com"`
	Password string `json:"password" binding:"required,min=5" example:"password123"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"johndoe@gmail.com"`
	Password string `json:"password" binding:"required,min=5" example:"password123"`
	Name     string `json:"name" binding:"required,min=3" example:"John Doe"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
