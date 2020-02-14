package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(n1, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(greeting("Vinicius"))
	fmt.Println(getSum(10, 125))
}
