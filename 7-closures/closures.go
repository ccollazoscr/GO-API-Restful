package main

import (
	"fmt"
)

func main() {
	fmt.Println(gorras(45, "$"))
}

func gorras(pedido float32, moneda string) (string, float32, string) {

	//Funcion anónima
	precio := func() float32 {
		return pedido * 7
	}

	return "El precio del pedido es: ", precio(), moneda
}
