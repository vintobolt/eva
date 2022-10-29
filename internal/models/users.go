package models

type User struct {
	Login    string
	Passwd   string
	Role     string
	Fullname string
}

type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
