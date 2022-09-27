package main

import "fmt"

func main() {
	// un map es una estructura, llave o valor.
	//para crear un mapa: palabra reservada map [llaves] [retorno o valores del tipo]
	m1 := make(map[string]int)

	m1["a"] = 8

	//imprime llave -> valor
	fmt.Println(m1)
	fmt.Println(m1["a"])
}
