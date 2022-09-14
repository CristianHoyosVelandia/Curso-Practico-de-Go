package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leerEntrada() string {
	//creo el objeto de newScanner
	scanner := bufio.NewScanner(os.Stdin)
	//print
	fmt.Println("Por favor ingrese la operacion a realizar: ")
	//input de go, scanner
	scanner.Scan()

	//guardo el valor scaneado
	return scanner.Text()

}

func leerOperador() string {
	//creo el objeto de newScanner
	scanner := bufio.NewScanner(os.Stdin)
	//print
	fmt.Println("Por favor ingrese el operador a realizar: ")
	//input de go, scanner
	scanner.Scan()

	//guardo el valor scaneado
	return scanner.Text()

}

// struct.
type calc struct{}

// receiver option
func (calc) operate(entrada string, operador string) int {

	valores := strings.Split(entrada, operador)

	// Cast valores from text to number AtoInt convierte a operador 1 o muestra el error
	operador1 := parsear(valores[0])
	operador2 := parsear(valores[1])

	switch operador {
	case "+":
		// fmt.Println("\n Suma de los dos operadores matematicamente: ", operador1+operador2)
		return operador1 + operador2
	case "-":
		// fmt.Println("\n Resta de los dos operadores matematicamente: ", operador1-operador2)
		return operador1 - operador2
	case "*":
		// fmt.Println("\n Multiplicacion de los dos operadores matematicamente: ", operador1*operador2)
		return operador1 * operador2
	case "/":
		// fmt.Println("\n Division de los dos operadores matematicamente: ", operador1/operador2)
		return operador1 / operador2
	default:
		fmt.Println("\n El operador", operador, " no esta soportado")
		return 0

	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func main() {

	entrada := leerEntrada()
	operador := leerOperador()
	//definicion de objeto de calc
	c := calc{}
	calculadora := c.operate(entrada, operador)
	fmt.Println("\n ", calculadora)

}
