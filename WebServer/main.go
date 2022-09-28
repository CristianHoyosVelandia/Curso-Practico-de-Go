package main

func main() {
	server := NewServer(":3000")

	//uso de la funcion handle con el path "/" y el handler
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	server.Listen()
}
