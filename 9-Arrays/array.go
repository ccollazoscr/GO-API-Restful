package main

import (
	"fmt"
)

func main() {
	var peliculas [3]string
	peliculas[0] = "La verdad duele"
	peliculas[1] = "Ciudadano ejemplar"
	peliculas[2] = "Gran torino"

	fmt.Println(peliculas)
	fmt.Println(peliculas[1])

	peliculas2 := [3]string{"La verdad duele",
		"CiudadanoEjemplar",
		"Batman"}
	fmt.Println(peliculas2)

	/*Array multidimencional*/
	var peliculas3 [3][2]string
	peliculas3[0][0] = "La verdad duele"
	peliculas3[0][1] = "Ciudadano ejemplar"
	peliculas3[1][0] = "Gran torino"
	peliculas3[1][1] = "El señor de los aniños"
	peliculas3[2][0] = "A todo gas"
	fmt.Println(peliculas3)
	fmt.Println(peliculas3[0])
	fmt.Println(peliculas3[0][1])

}
