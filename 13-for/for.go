package main

import "fmt"

func main() {
	tope := 26
	for i := 1; i <= tope; i++ {
		if i%2 == 0 {
			fmt.Println(i, "El numero es par")
		} else {
			fmt.Println(i, "El numero es impar")
		}
	}

	peliculas := []string{"Pelicula 1", "Pelicula 2", "Pelicula 3", "Pelicula 4", "Pelicula 5"}

	for i := 0; i < len(peliculas); i++ {
		if i%2 == 0 {
			fmt.Println("La pelicula "+peliculas[i]+" es par", i)
		} else {
			fmt.Println("La pelicula "+peliculas[i]+" es impar", i)
		}
	}

	//Simular foreach
	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}
}
