package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//creo el objeto de newScanner
	scanner := bufio.NewScanner(os.Stdin)
	//print
	fmt.Println("Ingrese la operacion (suma, de la forma numero + numero, ej: 2+2): ")
	//input de go, scanner
	scanner.Scan()

	//guardo el valor scaneado
	operacion := scanner.Text()

	//muestro la info guardada
	fmt.Println("\n La operacion ingresada es: ", operacion)

	valores := strings.Split(operacion, "+")

	fmt.Println("\n Estos son los valores ingresados: ", valores)
	fmt.Println("\n Primer y segundo valor sumados como texto: ", valores[0]+valores[1])

	// Cast valores from text to number AtoInt convierte a operador 1 o muestra el error
	operador1, err1 := strconv.Atoi(valores[0])

	if err1 != nil {
		fmt.Println("\n Hubo un error en operador 1")
	} else {
		fmt.Println("\n Continuamos con la operacion")
	}

	operador2, err2 := strconv.Atoi(valores[1])

	if err2 != nil {
		fmt.Println("\n Hubo un error en operador 2")
	} else {
		fmt.Println("\n Continuamos con la operacion")
	}

	fmt.Println("\n Suma de los dos operadores matematicamente: ", operador1+operador2)

}
