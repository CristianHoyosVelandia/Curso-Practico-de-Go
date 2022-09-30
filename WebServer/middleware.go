package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Auth")

			if flag == true {
				f(w, r)
			}
		}
	}
}

func Loggin() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			//defer designa una función a ejecutarse hasta el final de la ejecución del scope en el que se encuentre.
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			//esto permite realizar el salto al siguiente handlerfunciton
			hf(w, r)
		}
	}
}
