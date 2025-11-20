package main

import "fmt"

func main() {
	varMap := map[string]string{
		"Colombia":  "Bogotá",
		"Venezuela": "Caracas",
		"Brasil":    "Santiago",
		"Chile":     "Santiago",
	}

	fmt.Println("El mapa de paises: ", varMap)
	fmt.Println("La capital de venezuela es: ", varMap["Venezuela"])

	varMap["Argentina"] = "Buenos Aires"
	fmt.Println("El mapa de paises: ", varMap)

	delete(varMap, "Venezuela")
	fmt.Println("El mapa de paises: ", varMap)
	fmt.Println("La capital de venezuela es: ", varMap["Venezuela"])

}
