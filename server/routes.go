package server

import (
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", Index)

	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Getcontry(w, r)
		case http.MethodPost:
			Addcontry(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Solo se aceptan peticiones GET y POST")
		}
	})

}
