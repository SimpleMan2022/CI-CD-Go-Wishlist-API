package dto

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Data any `json:"data"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
