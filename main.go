package main

import "github.com/ponceca20/PRACTICAGO/server"

// importa modulo http

func main() {

	srv := server.New(":8080")
	//start the server whin control error
	error := srv.ListenAndServe()
	if error != nil {
		panic(error)
	}

}
