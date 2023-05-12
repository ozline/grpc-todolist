package types

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
