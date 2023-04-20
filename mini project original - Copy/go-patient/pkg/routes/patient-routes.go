package routes

import (
	//"github.com/Shrikantgiri25/go-patient/pkg/controllers"
	"github.com/Shrikantgiri25/go-patient/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPatientRoutes = func(router *mux.Router) {

	router.HandleFunc("/patients/", controllers.AddPatient).Methods("POST")

	router.HandleFunc("/patients/bookslot", controllers.BookSchedule).Methods("POST")

	router.HandleFunc("/patients/", controllers.GetPatients).Methods("GET")

	router.HandleFunc("/patient/{id}", controllers.GetPatientByID).Methods("GET")

	router.HandleFunc("/patient/{id}", controllers.DeletePatient).Methods("DELETE")

	///////////////////////////////////////////////////////////////////////////////
	router.HandleFunc("/patient/{pid}/doctor/{did}/{batch}/Booklsot", controllers.BookSlot).Methods("POST")

	router.HandleFunc("/getSameDocBatch/", controllers.GetSameDocBatch).Methods("GET") // both argue,memnts docid and batch

	// router.HandleFunc("/getSameDoc/", controllers.GetSameDoc).Methods("GET") //same doc

	// router.HandleFunc("/getSameBatch/", controllers.GetSameBatch).Methods("GET")
}
