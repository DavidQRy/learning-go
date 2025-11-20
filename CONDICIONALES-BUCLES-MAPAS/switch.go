package main

import "fmt"

func main() {
	var fruit string

	fmt.Println("Enter the name of a fruit: ")
	fmt.Scan(&fruit)

	switch fruit {
	case "apple":
		fmt.Println("You chose Apple")
	case "pear":
		fmt.Println("You chose Pear")
	case "grape":
		fmt.Println("You chose Grape")
	case "banana":
		fmt.Println("You chose Banana")
	case "strawberry":
		fmt.Println("You chose Strawberry")
	default:
		fmt.Println("I don't recognize that fruit")
	}
}
