package main

import "fmt"

func main() {
	fmt.Println("********* MI PROGRAMA CON GO *********")
	edad := 25

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
