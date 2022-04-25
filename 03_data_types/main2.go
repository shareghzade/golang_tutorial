package main

import "fmt"

func main() {
	// float32, float64 just gets floating numbers
	var imFloat32 float32
	// int8, int16, int32, int64 only gets number and it covers negetive numbers
	var imInt8 int8
	// uint8, uint16, uint32, uint64 only gets number and it doesn't cover the negetive numbers
	var imUint8 uint8

	imFloat32 = 13.4
	imInt8 = -2
	imUint8 = 2 // cant assign -2

	fmt.Printf("Float32: %v\nInt8: %v\nUint8: %v", imFloat32, imInt8, imUint8)

	// if you dont want to assign datatype to your variable
	// but you cant use this for const
	noAssign := "Hi i have auto datatype"
	fmt.Println("%v", noAssign)

}

/*

	for see all the integer datatype => https://www.tutorialandexample.com/integer-number-data-types-in-golang

*/
