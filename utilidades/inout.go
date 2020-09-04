package utilidades

import (
	"fmt"
	"strconv"
)

//Mensaje muestra tantos mensajes en pantalla como argumentos se pasan como parametro
func Mensaje(mensajes ...string) {
	for _, mensaje := range mensajes {
		fmt.Println(mensaje)
	}
}

//LeerOpcion determina la opcion seleccionada por el usuario
func LeerOpcion() (string, error) {
	var opcion string
	fmt.Scanln(&opcion)
	_, err := strconv.Atoi(opcion)
	return opcion, err
}

//LeerRegistro lee los datos de entrada para registrar un nuevo usuario en el sistema
func LeerRegistro() (string, string) {
	var usuario, password, passwordR string
	Mensaje("Ingrese usuario")
	fmt.Scanln(&usuario)
	Mensaje("Ingrese contraseña")
	fmt.Scanln(&password)
	for len(password) < 15 {
		Mensaje("Contraseña demasiado corta. Ingrese una nueva", "")
		fmt.Scanln(&password)
	}
	Mensaje("Repita contraseña")
	fmt.Scanln(&passwordR)
	for password != passwordR {
		Mensaje("Contraseñas no coinciden. Ingrese nuevamente", "")
		fmt.Scanln(&passwordR)
	}
	return usuario, password
}

//LeerInicio lee los datos que el usuario ingresa en pantalla para recuperar la sesion
func LeerInicio() (string, string) {
	var usuario, password string
	Mensaje("Ingrese usuario")
	fmt.Scanln(&usuario)
	Mensaje("Ingrese contraseña")
	fmt.Scanln(&password)
	return usuario, password
}

//CleanScreen limpia la terminal
func CleanScreen() {
	saltosLinea := 100
	for saltosLinea > 0 {
		Mensaje("")
		saltosLinea--
	}
}
