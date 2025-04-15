package dto

type RegisterRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Name           string `json:"name"`
	DepartmentName string `json:"department"`
}
