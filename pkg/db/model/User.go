package model

type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	GroupId      int64  `json:"group_id"`
}
