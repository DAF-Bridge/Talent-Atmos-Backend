package types

type JWT struct {
	Email  string `json:"email"`
	Exp    int64  `json:"exp"`
	UserID string `json:"user_id"`
}