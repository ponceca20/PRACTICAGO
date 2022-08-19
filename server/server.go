package server

import "net/http"

//crea struct country
type country struct {
	Name     string
	Language string
}

var countries []*country = []*country{}

//crea servidor
func New(addr string) *http.Server {
	InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
