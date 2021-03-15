package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// @Summary Gets "Hello World!"
// @Description Gets the string "Hello World!"
// @Response 200
// @Router /hello [get]
func getHello(w http.ResponseWriter, r *http.Request) {
	_hello(w, r, "World")
}

// @Summary Gets "Hello {name}!"
// @Description Gets the string "Hello {name}!"
// @Param name path string true "Name of the person to say hello to"
// @Success 200
// @Router /hello/{name} [get]
func getHelloWithName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	_hello(w, r, name)
}

func _hello(w http.ResponseWriter, r *http.Request, name string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	var response = fmt.Sprintf(`{"message": "Hello %v!"}`, name)
	w.Write([]byte(response))
}
