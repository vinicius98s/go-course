package main

import "fmt"

func main() {
	// // Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[0])

	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])
}
