package main

import (
	"fmt"

	"github.com/THEcanon001/AdminPw/logica"
	"github.com/THEcanon001/AdminPw/utilidades"
)

func main() {
	inicio()
}

func inicio() {
	utilidades.Mensaje("1. Iniciar Sesion", "2. Registrarse", "3. Eliminar Usuario", "4. Salir")
	opcion, err := utilidades.LeerOpcion()
	var intentos = 3
	for err != nil && intentos > 0 {
		intentos--
		fmt.Println("Intentos restantes", intentos)
		if intentos <= 0 {
			return
		}
		opcion, err = utilidades.LeerOpcion()
	}

	var resultado string
	switch opcion {
	case "1":
		utilidades.CleanScreen()
		utilidades.Mensaje("", "", "**************", "Inicio Usuario", "**************", "")
		err := logica.Inicio()
		if err != nil {
			reiniciar(resultado)
		} else {
			menu()
		}

	case "2":
		utilidades.CleanScreen()
		utilidades.Mensaje("", "", "****************", "Registro Usuario", "****************", "")
		resultado = logica.Registro()
		reiniciar(resultado)
	case "3":
		utilidades.CleanScreen()
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

func menu() {

}
