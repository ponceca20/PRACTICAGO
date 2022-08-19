package server

import "net/http"

//crea struct country
type country struct {
	Name     string
	language string
}

var countries []*country

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
