package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("lector")
	texto, errorDeFichero := ioutil.ReadFile("fichero2.txt")
	showError(errorDeFichero)
	fmt.Println(string(texto))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
