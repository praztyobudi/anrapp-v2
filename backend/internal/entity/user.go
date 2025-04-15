package entity

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Department *Department
}

type Department struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"departement"`
}
