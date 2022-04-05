package main

import (
	"fmt"
	"strconv"
)

func main() {
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Cccollazos "

	var apellidos string = "Ruiz"
	apellidos = "Collazos "

	//Definir variables utilizando otro método. Go infiere el tiepo de dato
	país := "Colombia"

	var prueba bool = true
	fmt.Println(prueba)

	var testDecimal float32 = 12.40
	fmt.Println(testDecimal)

	const year int = 2022 //Definiendo una constante
	fmt.Println("constante -> " + strconv.Itoa(year))

	fmt.Println("Hola mundo! - " + nombre + apellidos + país)
	fmt.Println(suma)
	fmt.Println(resta)
}
