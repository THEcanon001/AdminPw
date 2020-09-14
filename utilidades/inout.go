package utilidades

import (
	"fmt"
	"strconv"

	"github.com/THEcanon001/AdminPw/model"
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
func LeerRegistro() (string, string, string) {
	var usuario, password, passwordR, key32 string
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
	Mensaje("Ingrese una key que sera utilizada para decodificar sus contraseñas", "No olvide esta key",
		"Por seguridad no sera almacenada en el sistema")
	fmt.Scanln(&key32)
	for len(key32) != 32 {
		Mensaje("Key debe tener 1280 palabras")
		fmt.Scanln(&key32)
	}
	return usuario, password, key32
}

//LeerInicio lee los datos que el usuario ingresa en pantalla para recuperar la sesion
func LeerInicio() (string, string, string) {
	var usuario, password, key32 string
	Mensaje("Ingrese usuario")
	fmt.Scanln(&usuario)
	Mensaje("Ingrese contraseña")
	fmt.Scanln(&password)
	Mensaje("Ingrese key")
	fmt.Scanln(&key32)
	return usuario, password, key32
}

//LeerDatosEntrada lee los datos de la entrada para agregar un nuevo dato al usuario
func LeerDatosEntrada(ID int) (model.Data, error) {
	var dato, password string
	Mensaje("Ingrese nombre de dato")
	fmt.Scanln(&dato)
	Mensaje("Ingrese contraseña")
	fmt.Scanln(&password)
	d := model.Data{-1, dato, password, ID}
	return d, nil
}

//CleanScreen limpia la terminal
func CleanScreen() {
	saltosLinea := 100
	for saltosLinea > 0 {
		Mensaje("")
		saltosLinea--
	}
}

//PedirPassword pide la password al usuario para efectuar una accion posterior que implica datos
func PedirPassword() string {
	var password string
	Mensaje("Ingrese contraseña para continuar")
	fmt.Scanln(&password)
	return password
}

//Elegir escoge una opcion numerica
func Elegir() (string, error) {
	opcion, err := LeerOpcion()
	var intentos = 3
	for err != nil && intentos > 0 {
		intentos--
		fmt.Println("Intentos restantes", intentos)
		if intentos <= 0 {
			return "", err
		}
		opcion, err = LeerOpcion()
	}
	return opcion, nil
}
