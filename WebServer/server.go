// el archivo main va a poder leer todo lo que este en el archivo server
package main

import (
	"net/http"
)

type Server struct {
	//puerto del servidor para escuchar las conexiones
	port string

	//se agrega el atributo router que es un apuntador al struct Router de router.go
	router *Router
}

// Funcion tipo global para ser leida en otros archivos
// Sirve para instanciar el servidor y que sea capaz de escuchar las conexiones
// recive el puerto que tiene que estar escuchando y devuelve el servidor como tal
func NewServer(port string) *Server {
	return &Server{
		port: port,

		//router instanciado
		router: NewRouter(),
		//con esto el servidor ya es capaz de instanciar el router y de tenerlo como propiedad
	}
}

// Funcion tipo receiver, del struct Server, devuelve un error en caso de que haya problemas al conectar
func (s *Server) Listen() error {
	//el router va a ser el encargado de tomar las urls y procesarlas como se debe, crea el entry-point
	// los parametro son: el slash que es el punto de entrada, y el handler es el router recien creado
	http.Handle("/", s.router)

	//con la funcion listenanserve() del paquete http nos ayuda a escuchar todas las peticiones
	//colocas el puerto como primer parametro, el segundo es un handler
	//pero nosotros haremos nuestros handlers por eso se coloca nil
	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		return err
	}
	//si la ejecucion salio bien, retorna un valor nil
	return nil
}

// le decimos a go que no sabemos cuantos parametros vamos a enviar. cuando agregamos ...
// le decimos a GO que puede llegar 1 o mas middlewares
func (s *Server) AddMidleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

// Handle es el nombre de la ruta por ejemplo "/api" asignado a un handler especifico
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	//Asociacion del handler con la ruta, es decir, el mapa con la llave path asignado al handler
	//asi el servidor es capaz de agregar la ruta especifica a un handler especifico
	_, exist := s.router.rules[path]

	//aca realizamos la asociacion de existencia del mapa, dado que el mapa no existe, con el fin
	//de que el usuario conozca que evidentemente no existe, creamos el metodo y nos devolvera
	// un header de peticion diciendo que el methodo no esta defiinido para esa rutsa
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	//creo el handler con el metodo
	s.router.rules[path] = make(map[string]http.HandlerFunc)
	s.router.rules[path][method] = handler
}
