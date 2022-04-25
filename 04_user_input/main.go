package main

import "fmt"


func main() {
	// input for getting user Username
	fmt.Print("Enter Your Username: ")
	var userName string
	// "&userName" is a pointer
	fmt.Scan(&userName)

	// input for getting user Password
	fmt.Print("Enter Your Password: ")
	var pass int
	fmt.Scan(&pass)


	// Show the Username and password
	// NOTE: if you don't use integer in password input you will get "0" and it means nil

	fmt.Printf("\nUsername: %v\nPassword: %v", userName, pass)


}