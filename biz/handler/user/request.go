package user

type UserRegisterRequest struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	StatusCode int32
	StatusMSG  string
	UserID     int64
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
