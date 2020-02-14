package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method(value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Birthday: pointer receiver
func (p *Person) birthday() {
	p.age++
}

// Change name: pointer receiver
func (p *Person) changeName(newName string) {
	if p.age < 18 {
		return
	}

	p.firstName = newName
}

func main() {
	vinicius := Person{firstName: "Vinicius", lastName: "Sales", age: 21, city: "Diadema", gender: "M"}
	// vinicius := Person{"Vinicius", "Sales", "Diadema", "M", 21}

	fmt.Println(vinicius.age)
	vinicius.birthday()
	fmt.Println("Birthday")
	fmt.Println(vinicius.greet())

	vinicius.changeName("Diadema")
	fmt.Println(vinicius.firstName)
}
