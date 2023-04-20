package main

import (
	"log"
	"net/http"

	"github.com/Shrikantgiri25/go-doctor/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterDoctorRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":2010", r))
}
