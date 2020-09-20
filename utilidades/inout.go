package utilidades

import (
	"crypto/rand"
	"encoding/hex"
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
func LeerOpcion() (int, error) {
	var opcion string
	fmt.Scanln(&opcion)
	o, err := strconv.Atoi(opcion)
	return o, err
}

//LeerRegistro lee los datos de entrada para registrar un nuevo usuario en el sistema
func LeerRegistro() (string, string) {
	var usuario, password, passwordR, key32, mp string
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
	Mensaje("Ingrese o genere una key automática que sera utilizada para decodificar sus contraseñas. No olvide esta key", "Generar key automática? Y/N")
	fmt.Scanln(&mp)
	if mp == "y" || mp == "Y" {
		key32 = GenerarKey()
		fmt.Println("Su key es", key32)
	} else {
		Mensaje("Ingrese una key")
		fmt.Scanln(&key32)
		for len(key32) != 32 {
			Mensaje("Key debe tener 32 palabras")
			fmt.Scanln(&key32)
		}
	}

	return usuario, password
}

//LeerInicio lee los datos que el usuario ingresa en pantalla para recuperar la sesion
func LeerInicio() (string, string, string) {
	var usuario, password, key32 string
	Mensaje("Ingrese usuario")
	fmt.Scanln(&usuario)
	Mensaje("Ingrese contraseña")
	fmt.Scanln(&password)
	Mensaje("Ingrese su key")
	fmt.Scanln(&key32)
	return usuario, password, key32
}

//LeerModificacion lee los datos de entrada para registrar un nuevo usuario en el sistema
func LeerModificacion() (string, string) {
	var usuario, password, passwordR, mp string
	Mensaje("Ingrese usuario")
	fmt.Scanln(&usuario)
	Mensaje("Quiere modificar la contraseña para este usuario? Y/N")
	fmt.Scanln(&mp)
	if mp == "Y" || mp == "y" {
		Mensaje("Ingrese contraseña")
		fmt.Scanln(&password)
		for len(password) < 15 {
			Mensaje("Contraseña demasiado corta. Ingrese una nueva")
			fmt.Scanln(&password)
		}
		Mensaje("Repita contraseña")
		fmt.Scanln(&passwordR)
		for password != passwordR {
			Mensaje("Contraseñas no coinciden. Ingrese nuevamente")
			fmt.Scanln(&passwordR)
		}
	}
	return usuario, password
}

//LeerDatosEntrada lee los datos de la entrada para agregar un nuevo dato al usuario
func LeerDatosEntrada(ID int) (model.Data, error) {
	var name, user, password string
	Mensaje("Ingrese nombre")
	fmt.Scanln(&name)
	Mensaje("Ingrese usuario")
	fmt.Scanln(&user)
	Mensaje("Ingrese contraseña")
	fmt.Scanln(&password)
	d := model.Data{-1, name, user, password, ID}
	return d, nil
}

//PedirPassword pide la password al usuario para efectuar una accion posterior que implica datos
func PedirPassword() string {
	var password string
	Mensaje("Ingrese contraseña para continuar")
	fmt.Scanln(&password)
	return password
}

//Elegir escoge una opcion numerica
func Elegir() (int, error) {
	opcion, err := LeerOpcion()
	return opcion, err
}

//GenerarKey genera una key aleatoria
func GenerarKey() string {
	bytes := make([]byte, 32) //genero un random 32 byte key para AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key := hex.EncodeToString(bytes)
	return key //salt
}
