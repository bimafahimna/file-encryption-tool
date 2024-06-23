package encryption

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Encrypt() {
	file, err := os.Open("plain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text()) // token in unicode-char
		// fmt.Println(scanner.Bytes()) // token in bytes

	}
}

func EncryptFile(key []byte, filepath string) error {
	plaintext, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Use CBC mode with random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	cfb := cipher.NewCBCEncrypter(block, iv)
	var ciphertext []byte
	cfb.CryptBlocks(make([]byte, aes.BlockSize), plaintext)
	ciphertext = append(iv, ciphertext...)

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	return ioutil.WriteFile(filepath+".enc", []byte(encodedCiphertext), 0644)
}

func DecryptFile(key []byte, filepath string) error {
	ciphertext, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := decodedCiphertext[:aes.BlockSize]
	ciphertext = decodedCiphertext[aes.BlockSize:]

	cfb := cipher.NewCBCDecrypter(block, iv)
	var plaintext []byte
	cfb.CryptBlocks(make([]byte, aes.BlockSize), ciphertext)

	return ioutil.WriteFile(filepath+".dec", plaintext, 0644)
}
