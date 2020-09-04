package dao

import (
	"errors"
)

//ObtenerUsuario devuelve el hash y la salt del usuario asociado
func ObtenerUsuario(usuario string) (string, string) {
	return "hola", "chau"
}

//InsertarUsuario inserta un nuevo usuario
func InsertarUsuario(user string, pass string, salt string) error {
	return errors.New("can't work with 42")
}
