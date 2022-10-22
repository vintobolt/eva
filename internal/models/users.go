package models

type UserModel struct {
	Login    string
	Role     string
	Fullname string
}

type User struct {
	Login    string
	Passwd   string
	Role     string
	Fullname string
}
