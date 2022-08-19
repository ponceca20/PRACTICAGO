package server

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Solo se aceptan peticiones GET")
		return
	}
	fmt.Fprintf(w, "Hola mundo get")

}
func Getcontry(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "get contry get")
}
func Addcontry(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "get contry post")
}
