package main

import (
	"fmt"

	"github.com/THEcanon001/AdminPw/logica"
	"github.com/THEcanon001/AdminPw/model"
	"github.com/THEcanon001/AdminPw/utilidades"
)

func main() {
	inicio()
}

func inicio() {
	utilidades.Mensaje("**************", "Admin PW", "**************")
	utilidades.Mensaje("1. Iniciar Sesion", "2. Registrarse", "3. Modificar Usuario", "4. Eliminar Usuario", "5. Generar Secreto ", "6. Salir")
	o, err := utilidades.Elegir()
	if err != nil {
		return
	}
	var resultado string
	switch o {
	case 1:
		utilidades.Mensaje("**************", "Inicio Usuario", "**************")
		key, u, err := logica.Inicio()
		if err != nil {
			inicio()
		} else {
			menu(key, u)
		}

	case 2:
		utilidades.Mensaje("***********1*****", "Registro Usuario", "****************")
		resultado = logica.Registro()
		utilidades.Mensaje(resultado)
		inicio()
	case 3:
		utilidades.Mensaje("***********1*****", "Modificar Usuario", "****************")
		resultado = logica.Modificar()
		inicio()
		utilidades.Mensaje(resultado)
	case 4:
		resultado = logica.Eliminar()
		utilidades.Mensaje(resultado)
		inicio()
	case 5:
		resultado = utilidades.GenerarKey()
		utilidades.Mensaje(resultado)
		inicio()
	default:
		return
	}
}

func menu(key string, u model.Usuario) {
	utilidades.Mensaje("****************", "Menu Principal: "+u.User, "****************")
	utilidades.Mensaje("1. Agregar", "2. Ver", "3. Modificar", "4. Eliminar", "5. Cerrar Sesion", "6. Salir")
	opcion, err := utilidades.Elegir()
	if err != nil {
		return
	}
	switch opcion {
	case 1:
		utilidades.Mensaje("**************", "Agregar Dato", "**************")
		err := logica.AgregarDato(key, u)
		if err != nil {
			fmt.Println("error", err)
		} else {
			utilidades.Mensaje("Dato agregado de forma exitosa")
		}
		menu(key, u)
	case 2:
		utilidades.Mensaje("****************", "Ver Datos", "****************")
		logica.VerDatos(key, u)
		menu(key, u)
	case 3:
		utilidades.Mensaje("****************", "Modificar Dato", "****************")
		err := logica.ModificarDato(key, u)
		if err != nil {
			fmt.Println("error", err)
		} else {
			utilidades.Mensaje("Dato modificado de forma exitosa")
		}
		menu(key, u)
	case 4:
		utilidades.Mensaje("****************", "Eliminar Dato", "****************")
		err := logica.EliminarDato(key, u)
		if err != nil {
			fmt.Println("error", err)
		} else {
			utilidades.Mensaje("Dato eliminado de forma exitosa")
		}
		menu(key, u)
	case 5:
		inicio()
		break
	default:
		return
	}
}
