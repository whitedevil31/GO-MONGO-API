package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/whitedevil31/go-mongo-api/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterStudentRoutes(r)
	http.Handle("/", r)
	
	log.Fatal(http.ListenAndServe("localhost:5000", r))

}