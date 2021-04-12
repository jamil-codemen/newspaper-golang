package main

import (
	"fmt"
	"log"
	"net/http"
	"newspaper/controllers"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)


func main() {

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	
	router.Handle("/",staticC.Home).Methods("GET")
	router.Handle("/contact",staticC.Contact).Methods("GET")
	router.HandleFunc("/signup",usersC.New).Methods("GET")
	router.HandleFunc("/signup",usersC.Create).Methods("POST")

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT), router); err != nil {
		log.Fatal("error starting server: ", err)
	}
}

