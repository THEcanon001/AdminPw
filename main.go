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
	utilidades.Mensaje("1. Iniciar Sesion", "2. Registrarse", "3. Eliminar Usuario", "4. Salir")
	opcion, err := utilidades.Elegir()
	if err != nil {
		return
	}
	var resultado string
	utilidades.CleanScreen()
	switch opcion {
	case "1":
		utilidades.Mensaje("", "", "**************", "Inicio Usuario", "**************", "")
		key, u, err := logica.Inicio()
		if err != nil {
			reiniciar(resultado)
		} else {
			menu(key, u)
		}

	case "2":
		utilidades.Mensaje("", "", "****************", "Registro Usuario", "****************", "")
		resultado = logica.Registro()
		reiniciar(resultado)
	case "3":
		logica.Eliminar()
	default:
		return
	}
}

func reiniciar(resultado string) {
	utilidades.CleanScreen()
	utilidades.Mensaje(resultado)
	inicio()
}

func menu(key string, u model.Usuario) {
	utilidades.Mensaje("", "", "****************", "Menu Principal: "+u.Usuario, "****************", "")
	utilidades.Mensaje("1. Agregar", "2. Ver", "3. Modificar", "4. Eliminar", "5. Salir")
	opcion, err := utilidades.Elegir()
	if err != nil {
		return
	}
	utilidades.CleanScreen()
	switch opcion {
	case "1":
		utilidades.Mensaje("", "", "**************", "Agregar Dato", "**************", "")
		err := logica.AgregarDato(key, u)
		if err != nil {
			fmt.Println("error", err)
		} else {
			utilidades.Mensaje("Dato agregado de forma exitosa", "")
		}
	case "2":
		utilidades.Mensaje("", "", "****************", "Ver Datos", "****************", "")
		logica.VerDatos(key, u)
	case "3":
		utilidades.Mensaje("", "", "****************", "Modificar Dato", "****************", "")
		logica.ModificarDato(key, u)
	case "4":
		utilidades.Mensaje("", "", "****************", "Eliminar Dato", "****************", "")
		logica.EliminarDato(key, u)
	default:
		return
	}
	menu(key, u)
}


