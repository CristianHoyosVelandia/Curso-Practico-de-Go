package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()
	channel := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for i < 2 {
		for _, servidor := range servidores {
			go revisarServidores(servidor, channel)
		}

		for index := range servidores {
			fmt.Println(<-channel, index+1)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}
	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecuccion %q\n", tiempoPaso)
}

func revisarServidores(servidor string, channel chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		// fmt.Println(servidor, " No disponible :C")
		channel <- servidor + "No disponible :C"
	} else {
		// fmt.Println(servidor, " Funcionando correctamente :)")
		channel <- servidor + "Funcionando correctamente :)"
	}

}
