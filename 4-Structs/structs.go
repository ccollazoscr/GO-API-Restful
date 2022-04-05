package main

import (
	"fmt"
)

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	var gorra_negra = Gorra{
		marca:  "Nike",
		color:  "Negro",
		precio: 25.50,
		plana:  false}

	var gorra_blanca = Gorra{"Nike", "blanca", 25.50, true}

	fmt.Println(gorra_negra)
	fmt.Println(gorra_blanca.color)
}
