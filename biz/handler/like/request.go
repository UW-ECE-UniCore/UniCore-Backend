package like

type CheckLikeExistResponse struct {
	StatusCode int32  `json:"status_code"`
	Message    string `json:"message"`
	Existence  bool   `json:"existence"`
}

type ActionRequest struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

type ActionResponse struct {
	StatusCode int32  `json:"status_code"`
	Message    string `json:"message"`
	Existence  bool   `json:"existence"`
}
