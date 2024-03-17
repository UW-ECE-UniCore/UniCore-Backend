package user

type UserRegisterRequest struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	StatusCode int32  `json:"status_code"`
	Message    string `json:"message"`
}

type GetUserInfoResponse struct {
	Message   string `json:"message"`
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
	School    string `json:"school"`
	UserName  string `json:"username"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
}
