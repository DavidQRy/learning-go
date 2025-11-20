package main

import "fmt"

func main() {
	age := 17
	var permission bool = true

	/* if age >= 63 {
		fmt.Println("You can be retired.")
	} else if age >= 18 {
		fmt.Println("You are of legal age")
	} else {
		fmt.Println("You are a minor")
	}
	fmt.Println("End of program") */

	if age < 18 && permission {
		fmt.Println("The minor can travel alone because he has permission.")
	} else if age < 18 && !permission {
		fmt.Println("The minor cannot travel alone because he does not have permission.")
	} else {
		fmt.Println("The person can travel alone because they are of legal age")
	}
	fmt.Println("End of program")
}
