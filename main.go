package main

import "github.com/ponceca20/PRACTICAGO/server"

func main() {

	srv = server.New(":8080")
	//start the server whin control error
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
