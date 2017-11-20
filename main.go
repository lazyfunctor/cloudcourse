package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	bcrypt.GenerateFromPassword([]byte("test string !@#"), 11)
	fmt.Println("Hello, playground")
}
