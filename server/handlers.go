package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//funcion que sirve el index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")

	fmt.Fprintf(w, "Hola mundo pagina principal index")

}

//funcion que sirve contry
func Getcontry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

//funcion que agrega el contry
func Addcontry(w http.ResponseWriter, r *http.Request) {
	country := &country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error al leer el contry")
		return
	}
	countries = append(countries, country)
}
