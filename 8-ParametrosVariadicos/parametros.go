package main

import (
	"fmt"
)

func main() {
	pantalon("rojo", "largo", "sin bolsillos", "nike")
}

//Se reciben n o ningun parámetro tipo string
//Estas funciones se denomina vaiadicas
//https://www.digitalocean.com/community/tutorials/how-to-use-variadic-functions-in-go-esc
func pantalon(caracteristicas ...string) {
	for _, caracaracteristica := range caracteristicas {
		fmt.Println(caracaracteristica)
	}
}
