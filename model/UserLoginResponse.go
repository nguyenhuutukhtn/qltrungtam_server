package model

type UserLoginResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	RoleId   string `json:"role_id"`
}
