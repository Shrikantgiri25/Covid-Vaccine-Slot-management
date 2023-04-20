package routes

import (
	"github.com/Shrikantgiri25/go-doctor/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterDoctorRoutes = func(router *mux.Router) {

	router.HandleFunc("/doctors/", controllers.ScheduleDoctor).Methods("POST")

	router.HandleFunc("/doctors/", controllers.GetDoctors).Methods("GET")

	router.HandleFunc("/doctor/{id}", controllers.GetDoctorByID).Methods("GET")

	router.HandleFunc("/doctor/{id}", controllers.UpdateDoctor).Methods("PUT")

	router.HandleFunc("/doctor/{id}", controllers.DeleteDoctor).Methods("DELETE")

	router.HandleFunc("/bookSlot", controllers.BookSlot).Methods("POST")

	router.HandleFunc("/cancelSlot", controllers.CancelSlot).Methods(("POST"))

	/////////////////////////////////////////////////////////////////////////
	router.HandleFunc("/doctorName/{id}", controllers.DocName).Methods("GET")

}
