package entity

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Department *Department
}

type Department struct {
	ID         int    `json:"id"`
	Department string `json:"departement"`
}
