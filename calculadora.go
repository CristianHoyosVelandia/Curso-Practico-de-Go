package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculadora() string {
	//creo el objeto de newScanner
	scanner := bufio.NewScanner(os.Stdin)
	//print
	fmt.Println("Por favor ingrese la operacion a realizar: ")
	//input de go, scanner
	scanner.Scan()

	//guardo el valor scaneado
	return scanner.Text()

}

func main() {

	operacion := calculadora()

	//muestro la info guardada
	fmt.Println("\n La operacion ingresada es: ", operacion)
	operador := "-"
	valores := strings.Split(operacion, operador)

	fmt.Println("\n Estos son los valores ingresados: ", valores)
	// fmt.Println("\n Primer y segundo valor sumados como texto: ", valores[0]+valores[1])

	// Cast valores from text to number AtoInt convierte a operador 1 o muestra el error
	operador1, err1 := strconv.Atoi(valores[0])

	if err1 != nil {
		fmt.Println("\n Hubo un error en operador 1")
	} else {
		// fmt.Println("\n Continuamos con la operacion")
	}

	operador2, err2 := strconv.Atoi(valores[1])

	if err2 != nil {
		fmt.Println("\n Hubo un error en operador 2")
	} else {
		// fmt.Println("\n el operador dos esta correcto")
	}

	switch operador {
	case "+":
		fmt.Println("\n Suma de los dos operadores matematicamente: ", operador1+operador2)

	case "-":
		fmt.Println("\n Resta de los dos operadores matematicamente: ", operador1-operador2)

	case "*":
		fmt.Println("\n Multiplicacion de los dos operadores matematicamente: ", operador1*operador2)

	case "/":
		fmt.Println("\n Division de los dos operadores matematicamente: ", operador1/operador2)

	default:
		fmt.Println("\n El operador no esta soportado")
	}

}
