package main

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func HashPassword(password string, salt []byte) string {
	dk := pbkdf2.Key([]byte(password), salt, 4096, 32, sha1.New)
	return fmt.Sprintf("%x", dk)
}

func CheckPasswordHash(password string, salt []byte, hash string) bool {
	dk := pbkdf2.Key([]byte(password), salt, 4096, 32, sha1.New)
	return fmt.Sprintf("%x", dk) == hash
}

func main() {
	salt := make([]byte, 10)
	rand.Read(salt)

	password := "secret"
	hash := HashPassword(password, salt)

	fmt.Printf("Salt:     %x\n", salt)
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Hash:     %s\n", hash)

	match := CheckPasswordHash(password, salt, hash)
	fmt.Println("Match:   ", match)
}
