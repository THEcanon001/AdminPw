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

	"github.com/THEcanon001/AdminPw/model"
	"github.com/THEcanon001/AdminPw/utilidades"
	"golang.org/x/crypto/bcrypt"
)

const tam = 15

//Encrypt encripta un string con aes 256
func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	//Since the key is in string, we need to convert decode it to bytes
	hexstring := hex.EncodeToString([]byte(keyString))
	key, _ := hex.DecodeString(hexstring)
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
	hexstring := hex.EncodeToString([]byte(keyString))
	key, _ := hex.DecodeString(hexstring)
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

//EncriptarHash encripta la contraseña del usuario
func EncriptarHash(textoPlano string, key string) (string, string, error) {
	utilidades.Mensaje("Se esta almacenando al usuario de forma seguro, esto puede tardar unos minutos")
	salt := generarSalt()
	pepper := generarPepper(key)
	pwSaltPep := textoPlano + salt + pepper

	pwAsByte := []byte(pwSaltPep)
	hash, err := bcrypt.GenerateFromPassword(pwAsByte, bcrypt.DefaultCost) //DefaultCost es 10
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

func generarPepper(key string) string {
	ran, err := random(1, key)
	if err != nil {
		log.Fatalln(err)
	}
	return ran
}

func random(length int, key string) (string, error) {
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
func DesencriptarHash(u model.Usuario, textoPlano string) error {
	for i := 0; i < 256; i++ {
		pepper := string(i)
		pwSaltPep := textoPlano + u.Salt + pepper
		hashAsByte := []byte(u.Password)
		pwAsByte := []byte(pwSaltPep)
		error := bcrypt.CompareHashAndPassword(hashAsByte, pwAsByte)
		if error == nil {
			return nil
		}
	}
	return errors.New("Error al ingresar")
}
