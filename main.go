package main

import (
	"fmt"
	"math/rand"
)

func generatePassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[{]};:"
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

func main() {
	passwordLength := 12 // tamanho desejado da senha gerada
	password := generatePassword(passwordLength)
	fmt.Printf("Generated password: %s\n", password)
	fmt.Println()
}
