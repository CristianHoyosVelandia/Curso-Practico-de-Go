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

func (t *taskList) imprimirLista() {

	for _, lista := range t.task {
		fmt.Println(" Nombre de la tarea: ", lista.nombre)
		fmt.Println(" ¿Completado?: ", lista.completado)
	}

}

func (t *taskList) imprimirListaCompletados() {

	for _, lista := range t.task {
		if lista.completado == true {
			fmt.Println(" Nombre de la tarea: ", lista.nombre)
		}
	}

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

	t4 := &task{
		nombre:      "curso English",
		description: "completar el curso",
		completado:  false,
	}

	t5 := &task{
		nombre:      "curso C#",
		description: "completar el curso",
		completado:  false,
	}

	lista := &taskList{
		task: []*task{t, t2},
	}

	lista.agregaALista(t3)

	lista.imprimirLista()

	lista.task[0].marcarCompletada()

	fmt.Println("\n NUEVA LISTA DE TAREAS COMPLEATAS ")
	lista.imprimirListaCompletados()

	//maps llaves objetos
	mapaTareas := make(map[string]*taskList)

	mapaTareas["Cristian"] = lista

	lista2 := &taskList{
		task: []*task{t4, t5},
	}

	mapaTareas["OTO"] = lista2

	fmt.Println("Tareas de Cristian")
	mapaTareas["Cristian"].imprimirLista()
	fmt.Println("Tareas de OTO")
	mapaTareas["OTO"].imprimirLista()
}
