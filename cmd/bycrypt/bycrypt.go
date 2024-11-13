package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("invalid command: %v\n",os.Args[1])
	}
}

func hash(password string) {
	hashedBytes,err := bcrypt.GenerateFromPassword([]byte(password) ,bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing %q\n",password)
		return 
	}
	fmt.Println(string(hashedBytes))
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Printf("compare the password %q with the hash %q\n",password,hash)
	if err != nil {
		fmt.Println("password is invalid: %w\n",password)
		return 
	}
	fmt.Println("password is correct")
}