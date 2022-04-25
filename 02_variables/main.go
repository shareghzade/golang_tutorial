package main

import "fmt"

func main() {
	// For creating variable you should use "var" "variable name" "datatype"
	var name string = "Nick"
	var age int = 19
	// const is like var but you can't change the value
	const yourName = "Jony"

	fmt.Println("My Name is "+name+" My age is", age)

	// we use "%v" for variable
	fmt.Printf("Your name is %v", yourName)

}
