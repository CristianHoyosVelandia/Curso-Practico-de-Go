package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	cambiarValor(x)
	//imprimir la dirrecion hexadecimal de la localizacion de X en el espacio de memoria
	fmt.Println(&x)
	y := &x
	fmt.Println(y)
	fmt.Println(&y)
	//accede al valor que esta almacenado en la direccion hexadecimal en el espacio de memoria
	fmt.Println(*y)
}

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 36

}
