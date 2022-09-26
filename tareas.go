package main

import "fmt"

//FIRTS CRUD in GO, we use a new struct for arrays (List of objects)
type taskList struct {
	//[name]   [SLICE] [TYPE OF SLICE]
	task []*task
}

func (t *taskList) agregaALista(tl *task) {
	t.task = append(t.task, tl)
}

//eliminar de la lista
func (t *taskList) eliminarLista(index int) {
	//Los tres puntos al final de t.task[index + 1:]… (operador ellipsis)
	//es porque el segundo parámetro del append no es un slice y la función append recibe un item,
	//con este operador lo que hacemos es decirle a go que tome ese slice y
	//lo “desempaquete” para que sean muchos parámetros de 1 solo item y no un slice.

	t.task = append(t.task[:index], t.task[index+1:]...)
}

type task struct {
	nombre      string
	description string
	completado  bool
}

func (t *task) marcarCompletada() {
	t.completado = true
}

func (t *task) actualizarDescription(input string) {
	t.description = input
}

func (t *task) actualizarNombre(input string) {
	t.nombre = input
}

func main() {

	t := &task{
		nombre:      "completar mi curso de GO",
		description: "completar el curso de creacion de un servidor web",
		completado:  false,
	}

	t2 := &task{
		nombre:      "completar mi curso de Pyp",
		description: "completar el curso de creacion de un servidor web",
		completado:  false,
	}

	t3 := &task{
		nombre:      "completar mi curso de Java",
		description: "completar el curso de creacion de un servidor web",
		completado:  false,
	}

	lista := &taskList{
		task: []*task{t, t2},
	}

	lista.agregaALista(t3)
	//USAR CICLOS FOR
	for i := 0; i < len(lista.task); i++ {
		fmt.Println("Index: ", i, " Atributos: ", lista.task[i])
	}

	fmt.Println()
	//other For ways
	for index, tarea := range lista.task {
		fmt.Println("index: ", index, " Atributos: ", tarea)
	}

	//break in iterration cycle
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	//break in iterration cycle
	for i := 0; i < 4; i++ {
		//exception iteration
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// Size of a list
	// fmt.Println(len(lista.task))

	lista.eliminarLista(2)

}
