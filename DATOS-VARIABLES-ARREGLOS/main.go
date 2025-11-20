package main

import "fmt"

func main() {
	var nombrePersona string = "David"
	var apellidoPersona string = "Quintero"
	segundoApellido := "Ramírez"
	fmt.Println("Hola Mundo. Soy " + nombrePersona)
	fmt.Println("Mi Apellido es:  " + apellidoPersona)
	fmt.Println("Mi segundo Appellido es " + segundoApellido)

	/* Parte númerica */
	var asoActual int16 = 2025
	var asoReducido int8 = 100
	edad := 19

	fmt.Println("El Año actual es: ", asoActual)
	fmt.Println("El año abreviado es: ", asoReducido)
	fmt.Println("Mi edad es de ", edad)

	/* Arreglos */
	var listaFrutas = [5]string{"Pera", "Naranja"}
	fmt.Println(listaFrutas[4])

	/* 	listaPaises := [4]string{"", "Colombia", "Estados Unidos", "Argentina"}
	   	fmt.Println(listaPaises)
	   	listaPaises[0] = "Chile"
	   	fmt.Println(listaPaises) */

	listaPaises := []string{"Colombia", "Estados Unidos", "Argentina"}
	listaPaises = append(listaPaises, "Chile")
	fmt.Println(listaPaises)

	listaPaises2 := listaPaises[1:3]
	fmt.Println(listaPaises2)

	listaPaises3 := listaPaises[2:]
	fmt.Println(listaPaises3)

	listaPaises4 := listaPaises[:2]
	fmt.Println(listaPaises4)
}
