package main

import "fmt"

type Operation func(int, int) int

func sum(a int, b int) int      { return a + b }
func subtract(a int, b int) int { return a - b }
func multiply(a int, b int) int { return a * b }
func divide(a int, b int) int   { return a / b }

func presentResult(a int, b int, op Operation, name string) {
	fmt.Println(name, "=>", op(a, b))
}

func getOperation(op string) Operation {
	if op == "+" {
		return sum
	} else if op == "-" {
		return subtract
	} else if op == "*" {
		return multiply
	} else if op == "/" {
		return divide
	}
	return func(a, b int) int { return 0 }
}

func main() {
	presentResult(5, 6, sum, "Sum")
	presentResult(3, 9, subtract, "Subtraction")
	presentResult(2, 5, multiply, "Multiplication")
	presentResult(10, 2, divide, "Division")

	op := getOperation("*")
	fmt.Println("Using getOperation:", op(7, 3))

	func() {
		result := sum(4, 6)
		fmt.Println("Inner anonymous function result:", result)
	}()
}
