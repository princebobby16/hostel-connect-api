package main

import (
	"github.com/gorilla/mux"
	"hostelconnectserver/pkg"
	"log"
	"net/http"
)

const port  = ":8080"

func main()  {

	//Set up the router
	router := mux.NewRouter()

	//Set handle the handler
	smsHandler := http.HandlerFunc(pkg.SendEmail)

	//Set up the method and middleware for the handler
	router.Handle("/send/email", pkg.Middleware(smsHandler)).Methods(http.MethodPost)

	//Start the server
	log.Println("Server running on port ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
