package main

import "fmt"


func main() {
	var myArr = []string {}
	myArr = append(myArr, "Bmw", "Benz", "Peride", "206")
	for i:=0; i<2; i++ {
		// it will print the 2 first element
		// of "myArr"
		fmt.Println(myArr[i])
	}

	//////////////////

	var items = []int {}
	items = append(items, 1, 2, 10, 13)

	// if you dont want the index
	// you can put _ instead
	for index, i := range items {
		fmt.Println(index, i)
	}

}