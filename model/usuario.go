package model

//Usuario es un modelo de usuario en bd
type Usuario struct {
	ID       int
	User     string
	Password string
	Salt     string
}
