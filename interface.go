package main

import "fmt"

//how to define interface
type animal interface {
	//methods or atributes.
	mover() string
}

type perro struct{}

type pez struct{}

type pajaro struct{}

func (perro) mover() string {
	return "soy un perro y camino"
}

func (pez) mover() string {
	return "soy un pez y nado"
}

func (pajaro) mover() string {
	return "soy un pajaro y estoy volando"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	p := perro{}
	moverAnimal(p)
	pe := pez{}
	moverAnimal(pe)
	pa := pajaro{}
	moverAnimal(pa)
}
