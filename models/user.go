package models

type User struct {
	ID       int    `json:"id"`
	IIN      string `json:"iin"`
	Username string `json:"username"`
	Password string `json:"password"`
}
