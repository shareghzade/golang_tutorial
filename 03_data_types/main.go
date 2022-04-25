package main

import "fmt"

func main() {
	// You can define the variables and the assign them values
	// when you create variable like down below you have to specify the datatype
	var userName string
	var password int
	var isActive bool

	// assigning value to variables we created above
	userName = "Hamid"
	password = 123
	isActive = true

	// "\n" puts enter
	fmt.Println("\n------ User Info ------")
	fmt.Printf("\nUsername: %v\nPassword: %v\nisOnline: %v\n", userName, password, isActive)
	fmt.Println("-----------------------\n")

	// for geting datatype of a variable
	var myType = "I am String :)"
	fmt.Printf("DataType is: %T", myType)
}

// Data Types

/*


string => for text content
int => for integers
bool => for true and false



*/
