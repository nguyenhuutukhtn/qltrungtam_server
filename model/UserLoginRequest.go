package model

type UserLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
