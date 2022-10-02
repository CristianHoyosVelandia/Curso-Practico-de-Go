package main

import (
	"encoding/json"
	"net/http"
)

//Se definió un type Middleware que recibe y
//retorne http.HandlerFunc así todas las funciones que cuenten con esta firma se considerarán de tipo Middleware.

//Reciben y retornan lo mismo, para poder encadenarse y pasar el request entre varios middleware, hasta llegar al handler.
//Algo interesante de esta definición es como Middleware tiene un type func como parámetro,
//significa que recibirá una función como argumento que tiene una firma definida http.HandlerFunc.

// El operador elipsis (…) sirve para determinar una cantidad variable de parametros de un mismo tipo.
type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	//comilla simple invertida se saca con alt + 96
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) toJSON() ([]byte, error) {
	return json.Marshal(u)
}

type MetaData interface{}
