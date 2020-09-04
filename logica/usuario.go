package logica

//User es una estructura para representar los usuarios del sistema
type User struct {
	ID uint64
	U  string
	P  string
}

//Users contiene la lista de usuarios del sistema
var Users []User
