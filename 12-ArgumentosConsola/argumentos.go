package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("********* MI PROGRAMA CON GO *********")
	fmt.Println(os.Args)
	fmt.Println("Hola " + os.Args[1] + " bienvenido al programa en GO")
	edad, _ := strconv.Atoi(os.Args[2])

	if (edad >= 18 || edad == 17) && edad != 25 && edad != 99 {
		fmt.Println("Eres mayor de edad")
	} else if edad == 25 {
		fmt.Println("tienes 25")

	} else if edad == 99 {
		fmt.Println("Pronto moriras")
	} else {
		fmt.Println("Eres menor de edad")
	}

}
