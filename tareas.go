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

	fmt.Println(lista.task[0])

	lista.agregaALista(t3)

	fmt.Println(lista.task[2])
	// Size of a list
	fmt.Println(len(lista.task))

	lista.eliminarLista(2)

}
