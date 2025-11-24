package main

import (
	"fmt"
	"strings"
)

func main() {
	// cycle for
	/* var suma int = 0

	for i := 0; i <= 100; i++ {
		suma += i
	}
	fmt.Println(suma) */

	//cycle for each
	/* varMap := map[string]string{
		"Colombia":  "Bogotá",
		"Venezuela": "Caracas",
		"Brasil":    "Santiago",
		"Chile":     "Santiago",
	}
	for k, v := range varMap {
		fmt.Println("The capital of " + k + " is: " + v)
	} */

	// cycle do-while
	var fruit string
	for {
		println("enter a fruit")
		fmt.Scan(&fruit)
		fruit = strings.ToLower(fruit)
		if fruit == "orange" {
			fmt.Println("The fruit is an orange")
			break
		} else {
			fmt.Println("The fruit don't is orange")
		}
	}

}
