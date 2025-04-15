package dto

type UpdateUserRequest struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Name           string `json:"name"`
	DepartmentName string `json:"department"`
}
