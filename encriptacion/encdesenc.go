package encriptacion

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/THEcanon001/AdminPw/dao"
	"golang.org/x/crypto/bcrypt"
)

const key string = "e1274c1c13e3909ed4a4b23f6a1903e0f64a3b89b0e2440dbd5084e2f9fcb82c"
const tam = 15

//Encrypt encripta un string con aes 256
func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

//Decrypt desencripta un texto encriptado con aes
func Decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

//ObtenerKey devuelve la key generada para encriptar y desencriptar
func ObtenerKey() string {
	return key
}

//EncriptarHash encripta la contraseña del usuario
func EncriptarHash(textoPlano string) (string, string, error) {
	salt := generarSalt()
	pepper := generarPepper()
	pwSaltPep := textoPlano + salt + pepper

	pwAsByte := []byte(pwSaltPep)
	hash, err := bcrypt.GenerateFromPassword(pwAsByte, bcrypt.MaxCost) //DefaultCost es 10
	if err != nil {
		fmt.Println(err)
	}
	hashComoCadena := string(hash)
	return hashComoCadena, salt, err
}

func generarSalt() string {
	bytes := make([]byte, 32) //genero un random 32 byte key para AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key := hex.EncodeToString(bytes)
	return key //salt
}

func generarPepper() string {
	ran, err := random(1)
	if err != nil {
		log.Fatalln(err)
	}
	return ran
}

func random(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = key[b%byte(len(key))]
	}

	return string(bytes), nil
}

//DesencriptarHash verifica que el usuario sea el que dice ser
func DesencriptarHash(usuario string, textoPlano string) error {
	hash, salt := dao.ObtenerUsuario(usuario)
	for i := 0; i < 256; i++ {
		pepper := string(i)
		pwSaltPep := textoPlano + salt + pepper
		hashAsByte := []byte(hash)
		pwAsByte := []byte(pwSaltPep)
		error := bcrypt.CompareHashAndPassword(hashAsByte, pwAsByte)
		if error == nil {
			return nil
		}
	}
	return errors.New("Error al ingresar")
}
