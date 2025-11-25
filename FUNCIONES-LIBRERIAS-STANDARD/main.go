package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Name           string
	Qualifications []float64
}

func readCSV(fileName string) ([]Student, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var students []Student

	for _, record := range records {
		name := record[0]
		var qualifications []float64

		for _, qualificationStr := range record[1:] {
			qualification, err := strconv.ParseFloat(qualificationStr, 64)
			if err != nil {
				return nil, err
			}
			qualifications = append(qualifications, qualification)
		}

		student := Student{
			Name:           name,
			Qualifications: qualifications,
		}

		students = append(students, student)
	}

	return students, nil
}
func promQualifications(qualifications []float64) float64 {
	var sum float64

	for _, qualification := range qualifications {
		sum += qualification
	}
	return sum / float64(len(qualifications))
}

func writeCSVWithAverages(fileName string, students []Student) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	// Header
	writer.Write([]string{"Name", "Average"})

	for _, s := range students {
		avg := promQualifications(s.Qualifications)
		record := []string{s.Name, fmt.Sprintf("%.2f", avg)}
		writer.Write(record)
	}

	return nil
}

func main() {
	students, err := readCSV("notas.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = writeCSVWithAverages("averages.csv", students)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
	}

	for _, s := range students {
		fmt.Println("Name:", s.Name, "Qualifications:", s.Qualifications, "Average: ", promQualifications(s.Qualifications))
	}
}

/* type Operation func(int, int) int

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
*/
