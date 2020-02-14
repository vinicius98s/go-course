package main

import "fmt"

var name string = "Vinicius"

func main() {
	var age int32 = 21
	const cool = true
	secondName := "Sales"
	// size := 1.1

	email, phone := "vinicius.2010.s@gmail.com", 11959774831

	fmt.Println(name, secondName, age, email, phone, cool)
	fmt.Printf("%T\n", cool)
}
