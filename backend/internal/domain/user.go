package domain

type UUID string

type User struct {
	ID       UUID
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string
}
