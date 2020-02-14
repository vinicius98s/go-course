package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// // Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Jhon"] = "jhon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare assigning values
	emails := map[string]string{"Bob": "bob@gmail.com", "Jhon": "jhon@gmail.com", "Mike": "mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
	fmt.Println(len(emails))
}
