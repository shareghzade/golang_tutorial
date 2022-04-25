package main

import "fmt"

func main() {
	// slice has dynamic size

	var mySlice [] string
	mySlice = append(mySlice, "Hello", "Im", "Slice")

	fmt.Println(mySlice)
}