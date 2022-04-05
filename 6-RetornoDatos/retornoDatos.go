package main

import (
	"fmt"
)

func main() {
	fmt.Println(devolverText())
	fmt.Println(devolverDatos())
}

func devolverText() (string, string) {
	return "Hola", "Alumno"
}

/*Variables de salida*/
func devolverDatos() (dato1 string, dato2 int) {
	dato1 = "Ccollazos"
	dato2 = 35

	return
}
