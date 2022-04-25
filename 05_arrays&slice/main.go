package main

import "fmt"


func main() {
	// this "[2]" is your array siza "string" is your array type 
	// you can't put another datatype or put more than tow element in array 
	var myArray = [2] string{"Hamid", "Ali"}
	fmt.Println(myArray)
	// Print first element in array
	fmt.Println(myArray[0])

	// for pushing value to the array

	var myIntArray = [3] int {}
	myIntArray[0] = 2
	myIntArray[1] = 10
	myIntArray[2] = 30

	fmt.Println(myIntArray)
	// get the length of array
	fmt.Println("Array Length is:", len(myIntArray))


}