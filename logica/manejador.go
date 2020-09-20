package logica

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/THEcanon001/AdminPw/dao"
	"github.com/THEcanon001/AdminPw/encriptacion"
	"github.com/THEcanon001/AdminPw/model"
	"github.com/THEcanon001/AdminPw/utilidades"
)

//Registro agrega un nuevo usuario al sistema
func Registro() string {
	user, password := utilidades.LeerRegistro()
	passencript, salt, err := encriptacion.EncriptarHash(password)
	if err != nil {
		mostrarError(err, -1)
	}
	usuario := model.Usuario{-1, user, passencript, salt}
	err = dao.InsertarUsuario(usuario)
	if err != nil {
		mostrarError(err, -1)
	}
	return "Registro de usuario exitoso"
}

//Inicio inicia sesion de un usuario existente
func Inicio() (string, model.Usuario, error) {
	usuario, password, key := utilidades.LeerInicio()
	u, err := dao.ObtenerUsuario(usuario)
	if err != nil {
		mostrarError(err, -1)
	}
	err = encriptacion.DesencriptarHash(u, password)
	if err != nil {
		mostrarError(err, 1)
	}
	return key, u, err
}

//Modificar modifica un usuario existente en el sistema
func Modificar() string {
	var usuario string
	fmt.Println("Ingrese el nombre de usuario que desea modificar.")
	fmt.Scanln(&usuario)

	u, err := dao.ObtenerUsuario(usuario)
	if err != nil || u.ID == 0 {
		mostrarError(err, -1)
	}
	user, password := utilidades.LeerModificacion()
	pedirPasswordUsuario(u)
	if user != "" && user != "\n" {
		u.User = user
	}

	if password != "" && password != "\n" {
		passencript, salt, err := encriptacion.EncriptarHash(password)
		if err != nil {
			mostrarError(err, -1)
		}
		u.Password = passencript
		u.Salt = salt
	}

	err = dao.ModificarUsuario(u)
	if err != nil {
		mostrarError(err, -1)
	}
	return "Modificacion de usuario exitosa"
}

//Eliminar elimina un usuario del sistema y sus contraseñas almacenadas
func Eliminar() string {
	var usuario, mp string
	fmt.Println("Ingrese el nombre de usuario que desea eliminar.")
	fmt.Scanln(&usuario)

	u, err := dao.ObtenerUsuario(usuario)
	if err != nil || u.ID == 0 {
		mostrarError(err, -1)
	}
	fmt.Println("Esta seguro que desea eliminar el usuario" + usuario + "? Se borraran todas las contraseñas almacenadas Y/N")
	fmt.Scanln(&mp)
	if mp == "y" || mp == "Y" {
		pedirPasswordUsuario(u)
		err = dao.EliminarUsuario(u)
		if err != nil {
			mostrarError(err, -1)
		}
		return "Eliminacion de usuario exitosa"
	} else {
		return "Operacion cancelada"
	}
}

//AgregarDato agrega informacion al usuario
func AgregarDato(key string, u model.Usuario) error {
	d, err := utilidades.LeerDatosEntrada(u.ID)
	if err != nil {
		mostrarError(err, -1)
	}
	pedirPasswordUsuario(u)
	d.Password = encriptacion.Encrypt(d.Password, key)
	err = dao.AgregarDato(u.ID, d)
	if err != nil {
		mostrarError(err, -1)
	}
	return nil
}

//VerDatos muestra la informacion de todas las claves almacenadas
func VerDatos(key string, u model.Usuario) {
	fmt.Println("Que contraseña desea ver?")
	r := mostrarDatos(key, u)
	verDatosYContinuar(key, u, r)
}

func verDatosYContinuar(key string, um model.Usuario, r map[int]model.Data) {
	os, err := utilidades.Elegir()
	if err != nil || os > len(r) || os < 1 {
		return
	}
	pedirPasswordUsuario(um)
	p := encriptacion.Decrypt(r[os].Password, key)
	fmt.Println(r[os].Name + "|" + r[os].User + "|" + p)
	verDatosYContinuar(key, um, r)
}

//ModificarDato modifica un dato existente del usuario
func ModificarDato(key string, u model.Usuario) error {
	fmt.Println("Que dato desea modificar?")
	r := mostrarDatos(key, u)
	os, err := utilidades.Elegir()
	if err != nil || os > len(r) || os < 1 {
		return errors.New("No se ingreso ninguna opcion.")
	}

	d, err := utilidades.LeerDatosEntrada(u.ID)
	if err != nil {
		mostrarError(err, -1)
	}
	pedirPasswordUsuario(u)
	d.IDusuario = u.ID
	d.ID = os
	fmt.Println("nombre", d.Name)
	fmt.Println("dato", d.User)
	fmt.Println("p", d.Password)
	if d.Name == "" || d.Name == "\n" {
		d.Name = r[os].Name
	}

	if d.User == "" || d.User == "\n" {
		d.User = r[os].User
	}

	if d.Password != "" && d.Password != "\n" && len(d.Password) > 1 {
		d.Password = encriptacion.Encrypt(d.Password, key)
	} else {
		d.Password = r[os].Password
	}
	fmt.Println("nombre", d.Name)
	fmt.Println("dato", d.User)
	fmt.Println("p", d.Password)
	err = dao.ModificarDato(u.ID, d)
	if err != nil {
		mostrarError(err, 1)
	}
	return nil
}

//EliminarDato elimina un dato existente del usuario
func EliminarDato(key string, u model.Usuario) error {
	fmt.Println("Que dato desea eliminar?")
	r := mostrarDatos(key, u)
	os, err := utilidades.Elegir()
	if err != nil || os > len(r) || os < 1 {
		return errors.New("No se ingreso ninguna opcion.")
	}
	pedirPasswordUsuario(u)
	err = dao.EliminarDato(u.ID, os)
	if err != nil {
		mostrarError(err, 1)
	}
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

func mostrarDatos(key string, u model.Usuario) map[int]model.Data {
	r, err := dao.VerDatos(u.ID)
	if err != nil {
		mostrarError(err, -1)
	}
	keys := make([]int, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(strconv.Itoa(k)+".", r[k].Name, "|", r[k].User)
	}
	fmt.Println("s.", "Volver")
	return r
}

func pedirPasswordUsuario(u model.Usuario) {
	p1 := utilidades.PedirPassword()
	err := encriptacion.DesencriptarHash(u, p1)
	if err != nil {
		mostrarError(err, 1)
	}
}
