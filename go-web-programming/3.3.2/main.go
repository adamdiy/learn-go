package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}


func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server {
		Addr: "127.0.0.1:8000",
	}

	http.Handle("/Hello", &hello)
	http.Handle("/World", &world)

	server.ListenAndServe()
}
