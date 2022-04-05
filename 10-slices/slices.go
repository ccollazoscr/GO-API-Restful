package main

import (
	"fmt"
)

func main() {
	//Slice es un array pero no tiene el numero de indice marcado

	peliculas := []string{"La verdad duele", "Ciudadano ejemplar", "Gran torino", "Superman"}
	peliculas = append(peliculas, "Sin limites")

	fmt.Println(peliculas)

	//Tamaño del array
	fmt.Println(len(peliculas))

	//Obtener elementos del 0 al 3
	fmt.Println(peliculas[0:3])
}
