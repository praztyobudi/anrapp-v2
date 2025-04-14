package entity

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Department string `json:"department"`
}
