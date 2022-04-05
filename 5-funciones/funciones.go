package main

import (
	"fmt"
)

func main() {
	var numero1 int = 10
	var numero2 int = 6

	calculadora(numero1, numero2)
	holaMundo()
}

func holaMundo() {
	fmt.Println("Hola mundo!!")
}

func operacion(n1 int, n2 int, op string) int {
	var resultado int
	if op == "+" {
		resultado = n1 + n2
	}

	if op == "-" {
		resultado = n1 - n2
	}

	if op == "*" {
		resultado = n1 * n2
	}

	if op == "/" {
		resultado = n1 * n2
	}

	if op == "%" {
		resultado = n1 % n2
	}
	return resultado
}

func calculadora(numero1 int, numero2 int) {
	//Suma
	fmt.Print("La suma es: ")
	fmt.Println(operacion(numero1, numero2, "+"))

	//Resta
	fmt.Print("La resta es: ")
	fmt.Println(operacion(numero1, numero2, "-"))

	//Multiplicacion
	fmt.Print("La multiplicacion es: ")
	fmt.Println(operacion(numero1, numero2, "*"))

	//División
	fmt.Print("La división es: ")
	fmt.Println(operacion(numero1, numero2, "/"))

	//Resto división
	fmt.Print("El resto de la división es: ")
	fmt.Println(operacion(numero1, numero2, "%"))

}
