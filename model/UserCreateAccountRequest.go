package model

type UserCreateAccountRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
	Role     string `json:""role,omitempty`
}
