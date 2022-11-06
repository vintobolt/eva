package models

type User struct {
	ID       string
	Login    string
	Passwd   string
	Role     string
	Fullname string
	Active   bool
}

type SignUp struct {
	Username string `json:"username" validation="required"`
	Password string `json:"password" validation="required"`
	Fullname string `json:"fullname" validation="required"`
}

type SignIn struct {
	Username string `json:"username" validation="required"`
	Password string `json:"password" validation="required"`
}
