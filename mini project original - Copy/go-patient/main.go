package main

import (
	"log"
	"net/http"

	"github.com/Shrikantgiri25/go-patient/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterPatientRoutes(r)

	//Handle registers the handler for the given pattern in the DefaultServeMux.
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3010", r))

}
