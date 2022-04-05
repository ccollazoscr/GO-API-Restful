package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("lector")

	nuevo_texto := os.Args[1] + " \n"
	//escribir := ioutil.WriteFile("fichero.txt", nuevo_texto, 0777) //-> Crea un nuevo archivo
	fichero, err := os.OpenFile("fichero.txt", os.O_APPEND, 0777)
	showError(err)

	escribir, err := fichero.WriteString(nuevo_texto)
	fmt.Println(escribir)
	showError(err)
	fichero.Close()

	texto, errorDeFichero := ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)
	fmt.Println(string(texto))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
