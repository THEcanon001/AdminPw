package logica

import (
	"fmt"
	"log"
	"strconv"

	"github.com/THEcanon001/AdminPw/dao"
	"github.com/THEcanon001/AdminPw/encriptacion"
	"github.com/THEcanon001/AdminPw/model"
	"github.com/THEcanon001/AdminPw/utilidades"
)

//Registro agrega un nuevo usuario al sistema
func Registro() string {
	user, password, key := utilidades.LeerRegistro()
	passencript, salt, err := encriptacion.EncriptarHash(password, key)
	if err != nil {
		mostrarError(err, -1)
	}
	usuario := model.Usuario{-1, user, passencript, salt}
	err = dao.InsertarUsuario(usuario, key)
	if err != nil {
		mostrarError(err, -1)
	}
	return "Registro Exitoso"
}

//Inicio inicia sesion de un usuario existente
func Inicio() (string, model.Usuario, error) {
	usuario, password, key := utilidades.LeerInicio()
	u, err := dao.ObtenerUsuario(usuario, key)
	if err != nil {
		mostrarError(err, -1)
	}
	err = encriptacion.DesencriptarHash(u, password)
	if err != nil {
		mostrarError(err, 1)
	}
	return key, u, err
}

//Eliminar elimina un usuario del sistema y sus contraseñas almacenadas
func Eliminar() {
	utilidades.Mensaje("usuario eliminado")
}

//AgregarDato agrega informacion al usuario
func AgregarDato(key string, u model.Usuario) error {
	d, err := utilidades.LeerDatosEntrada(u.ID)
	if err != nil {
		mostrarError(err, -1)
	}
	p := utilidades.PedirPassword()
	err = encriptacion.DesencriptarHash(u, p)
	if err != nil {
		mostrarError(err, 1)
	}
	d.Password = encriptacion.Encrypt(d.Password, key)
	err = dao.AgregarDato(u.ID, d, key)
	if err != nil {
		mostrarError(err, -1)
	}
	return nil
}

//VerDatos muestra la informacion de todas las claves almacenadas
func VerDatos(key string, u model.Usuario) {
	r, err := dao.VerDatos(u.ID, key)
	if err != nil {
		mostrarError(err, -1)
	}
	fmt.Println("Que contraseña desea ver?")
	for k, v := range r {
		fmt.Println(k, ".", v.Nombre)
	}
	fmt.Println(len(r)+1, ".", "Salir")
	os, err := utilidades.Elegir()
	if err != nil {
		return
	}

	o, _ := strconv.Atoi(os)
	if o == len(r)+1 {
		return
	}
	p1 := utilidades.PedirPassword()
	err = encriptacion.DesencriptarHash(u, p1)
	if err != nil {
		mostrarError(err, 1)
	}
	p := encriptacion.Decrypt(r[o].Password, key)
	verDatosYContinuar(r[o].Nombre, p, key, u)
}

func verDatosYContinuar(n string, p string, key string, u model.Usuario) {
	utilidades.CleanScreen()
	utilidades.Mensaje("**************", n+"->"+p, "**************", "")
	VerDatos(key, u)
}

//ModificarDato modifica un dato existente del usuario
func ModificarDato(key string, u model.Usuario) error {
	return nil
}

//EliminarDato elimina un dato existente del usuario
func EliminarDato(key string, u model.Usuario) error {
	return nil
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
