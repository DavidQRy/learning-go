package main

func sum(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func presentResult(a int, b int, operation string) {
	if operation == "+" {
		println("The sum of ", a, " + ", b, " is: ", sum(a, b))
	} else if operation == "-" {
		println("The subtraction of ", a, " - ", b, " is: ", subtract(a, b))
	} else if operation == "*" {
		println("The multiplication of ", a, " * ", b, " is: ", multiply(a, b))
	} else if operation == "/" {
		println("The division of ", a, " / ", b, " is: ", divide(a, b))
	} else {
		println("Invalid operation")
	}
}

func main() {
	presentResult(5, 6, "+")
	presentResult(3, 9, "-")
	presentResult(2, 5, "*")
	presentResult(10, 2, "/")
}
