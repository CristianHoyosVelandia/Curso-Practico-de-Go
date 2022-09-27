package main

import "fmt"

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) caminar() string {
	return "soy un perro y camino"
}

func (pez) nada() string {
	return "soy un pez y nado"
}

func (pajaro) Vuela() string {
	return "soy un pajaro y estoy volando"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(pe pez) {
	fmt.Println(pe.nada())
}

func moverPajaro(pa pajaro) {
	fmt.Println(pa.Vuela())
}

func main() {
	p := perro{}
	moverPerro(p)
	pe := pez{}
	moverPez(pe)
	pa := pajaro{}
	moverPajaro(pa)

}
