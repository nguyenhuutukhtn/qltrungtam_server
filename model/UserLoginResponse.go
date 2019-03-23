package model

type UserLoginResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Id       string `json:"id"`
	Role     string `json:"role"`
	RoleId   string `json:"role_id"`
}
