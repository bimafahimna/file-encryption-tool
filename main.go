package main

import (
	"file-encryption-tool/encryption"
	"fmt"
)

func main() {
	// encryption.Encrypt()
	// Replace with your desired password
	password := "SECRET"
	key := []byte(password)

	// Specify the file to encrypt/decrypt
	filepath := "plain.txt"

	// Choose between encryption or decryption
	action := "encrypt" // or "decrypt"

	if action == "encrypt" {
		err := encryption.EncryptFile(key, filepath)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}
		fmt.Println("File encrypted successfully!")
	} else if action == "decrypt" {
		err := encryption.DecryptFile(key, filepath+".enc")
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
		fmt.Println("File decrypted successfully!")
	} else {
		fmt.Println("Invalid action. Please choose 'encrypt' or 'decrypt'.")
	}
}
