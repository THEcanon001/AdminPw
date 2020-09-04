package logica

import (
	"fmt"
	"log"

	"github.com/THEcanon001/AdminPw/dao"
	"github.com/THEcanon001/AdminPw/encriptacion"
	"github.com/THEcanon001/AdminPw/utilidades"
)

//Registro agrega un nuevo usuario al sistema
func Registro() string {
	user, password := utilidades.LeerRegistro()
	passencript, salt, err := encriptacion.EncriptarHash(password)
	if err != nil {
		mostrarError(err, -1)
	}
	err = dao.InsertarUsuario(user, passencript, salt)
	if err != nil {
		mostrarError(err, -1)
	}
	return "Registro"
}

//Inicio inicia sesion de un usuario existente
func Inicio() error {
	usuario, password := utilidades.LeerInicio()
	err := encriptacion.DesencriptarHash(usuario, password)
	if err != nil {
		mostrarError(err, 1)
	}
	return err
}

//Eliminar elimina un usuario del sistema y sus contraseñas almacenadas
func Eliminar() {
	utilidades.Mensaje("usuario eliminado")
}

/*
	Error -1 = abortar
	Error 1 = reintento
*/
func mostrarError(err error, estado int) {
	switch estado {
	case 1:
		log.Fatal(err)
		utilidades.Mensaje("Usuario y/o contraseña incorrectos")
	default:
		var salir string
		log.Fatal(err)
		utilidades.Mensaje("Programa finalizado. Pulse una tecla para salir.")
		fmt.Scanln(&salir)
		return
	}

}
