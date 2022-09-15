package main // package name

// import libraries necessaries for executing the file
import "fmt"

// function main is the start point for all code.
func main() {

	//define variables
	//remenber
	// [var] [variables name] [type] = ["inicialized"]
	var mensaje string = "Hola mundo"
	//other form , go infiered the type of data
	mensajeFacil := "hola mundo usando :="

	fmt.Println(mensaje)
	fmt.Println(mensajeFacil)

	// float numbers
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)

	//integer numbers
	var c int = 10
	d := 3
	fmt.Println(c / d)

	// boolean
	var x bool = true
	y := false
	// OR logic
	fmt.Println(x || y)
	// AND logic
	fmt.Println(x && y)
	// NOT logic
	fmt.Println(!x)

	privada()

	Publica()

	printHelloWorld()
}

// func types
func privada() {
	fmt.Println("Ejecutar lógica que no necesita ser exportada (pertenece solo a este paquete)")
}

func Publica() {
	fmt.Println("Lógica que quiero exportar a otros paquetes")
}

// defer
func printHelloWorld() {
	defer fmt.Println("World!")
	fmt.Println("Hello")
}
