package models

type UserModel struct {
	Login    string
	Role     string
	Fullname string
}

type UserCredsModel struct {
	Login    string
	Password string
}

type UserRole struct {
	Login string
	Role  string
}
