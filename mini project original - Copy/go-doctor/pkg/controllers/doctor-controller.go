package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Shrikantgiri25/go-doctor/pkg/models"
	"github.com/Shrikantgiri25/go-doctor/pkg/utils"
	"github.com/gorilla/mux"
)

var NewDoctor models.Doctor

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	newDoctors := models.GetAllDoctors()
	res, _ := json.Marshal(newDoctors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorId := vars["id"]
	ID, err := strconv.ParseInt(doctorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	doctorDetails, _ := models.GetDoctorsById(ID)
	if (doctorDetails.Morning_count == 0) && (doctorDetails.Afternoon_count == 0) && (doctorDetails.Evening_count == 0) && doctorDetails.Doctor_name == "" {
		message := "Error not found"
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
	} else {
		res, _ := json.Marshal(doctorDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
func ScheduleDoctor(w http.ResponseWriter, r *http.Request) {
	addDoctor := &models.Doctor{}
	utils.ParseBody(r, addDoctor)

	if (addDoctor.Morning_count == 4 || addDoctor.Morning_count == 0) && (addDoctor.Afternoon_count == 4 || addDoctor.Afternoon_count == 0) && (addDoctor.Evening_count == 4 || addDoctor.Evening_count == 0) {
		b := addDoctor.ScheduleDoctor()
		res, _ := json.Marshal(b)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		message := "You can schedule Batch with size 4 or no Batch(Count as zero)"
		res, _ := json.Marshal(message)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(res)
	}

}

func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorId := vars["id"]
	ID, err := strconv.ParseInt(doctorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	doctorDetails, _ := models.GetDoctorsById(ID)
	if (doctorDetails.Morning_count > 0 && doctorDetails.Morning_count < 4) || (doctorDetails.Afternoon_count > 0 && doctorDetails.Afternoon_count < 4) || (doctorDetails.Evening_count > 0 && doctorDetails.Evening_count < 4) {
		message := "Can't Delete becuase schedules are Booked"
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(res)
	} else if (doctorDetails.Morning_count == 0) && (doctorDetails.Afternoon_count == 0) && (doctorDetails.Evening_count == 0) && doctorDetails.Doctor_name == "" {
		message := "Error not found"
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
	} else {
		_ = models.DeleteDoctor(ID)
		message := "Deleted Successfully"
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	var updateDoctor = &models.Doctor{}
	utils.ParseBody(r, updateDoctor)

	vars := mux.Vars(r)

	doctorId := vars["id"]
	ID, err := strconv.ParseInt(doctorId, 0, 0)
	if err != nil {
		fmt.Println("Error while string parsing")
	}
	doctorDetails, db := models.GetDoctorsById(ID)
	//morning
	if updateDoctor.Morning_count != doctorDetails.Morning_count {
		if doctorDetails.Morning_count == 0 || doctorDetails.Morning_count == 4 {
			doctorDetails.Morning_count = updateDoctor.Morning_count
			doctorDetails.Doctor_name = updateDoctor.Doctor_name
			updateDoctor.Afternoon_count = doctorDetails.Afternoon_count
			updateDoctor.Evening_count = doctorDetails.Evening_count
			doctorDetails = updateDoctor
		} else {
			message := "Can't update. Some patients are scheduled in morning batch"
			res, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(res)
		}
	}

	// Afternoon
	if updateDoctor.Afternoon_count != doctorDetails.Afternoon_count {
		if doctorDetails.Afternoon_count == 0 || doctorDetails.Afternoon_count == 4 {
			doctorDetails.Afternoon_count = updateDoctor.Afternoon_count
			doctorDetails.Doctor_name = updateDoctor.Doctor_name
			updateDoctor.Morning_count = doctorDetails.Morning_count
			updateDoctor.Evening_count = doctorDetails.Evening_count
			doctorDetails = updateDoctor
		} else {
			message := "Can't update. Some patients are scheduled in afternoon batch"
			res, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(res)
		}
	}
	// evening
	if updateDoctor.Evening_count != doctorDetails.Evening_count {
		if doctorDetails.Evening_count == 0 || doctorDetails.Evening_count == 4 {
			doctorDetails.Evening_count = updateDoctor.Evening_count
			doctorDetails.Doctor_name = updateDoctor.Doctor_name
			updateDoctor.Afternoon_count = doctorDetails.Afternoon_count
			updateDoctor.Morning_count = doctorDetails.Morning_count
			doctorDetails = updateDoctor
		} else {
			message := "Can't update. Some patients are scheduled inf morning batch"
			res, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(res)
		}
	}
	db.Save(&doctorDetails)
	res, _ := json.Marshal(doctorDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func BookSlot(w http.ResponseWriter, r *http.Request) {
	doctorToBook := &models.DoctorToBook{}

	utils.ParseBody(r, doctorToBook)

	doctorID := doctorToBook.DoctorId
	ID, err := strconv.ParseInt(doctorID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	doctorDetails, db := models.GetDoctorsById(ID)

	if doctorToBook.Batch == "Morning" {
		if doctorDetails.Morning_count > 0 {
			doctorDetails.Morning_count--
		} else {
			message := "Can't its already full"
			res, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(res)
		}
	} else if doctorToBook.Batch == "Afternoon" {
		if doctorDetails.Afternoon_count > 0 {
			doctorDetails.Afternoon_count--
		} else {
			message := "Can't its already full"
			res, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(res)
		}
	} else if doctorToBook.Batch == "Evening" {
		if doctorDetails.Evening_count > 0 {
			doctorDetails.Evening_count--
		} else {
			message := "Can't its already full"
			res, _ := json.Marshal(message)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(res)
		}
	}
	db.Save(&doctorDetails)
	res, _ := json.Marshal(doctorDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func CancelSlot(w http.ResponseWriter, r *http.Request) {
	doctorToCancel := &models.DoctorToBook{}

	utils.ParseBody(r, doctorToCancel)

	doctorId := doctorToCancel.DoctorId

	ID, err := strconv.ParseInt(doctorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	doctorDetails, db := models.GetDoctorsById(ID)
	if doctorToCancel.Batch == "Morning" {
		doctorDetails.Morning_count++
	} else if doctorToCancel.Batch == "Afternoon" {
		doctorDetails.Afternoon_count++
	} else if doctorToCancel.Batch == "Evening" {
		doctorDetails.Evening_count++
	}

	db.Save(&doctorDetails)
	res, _ := json.Marshal(doctorDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotAcceptable)
	w.Write(res)

}
func DocName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorId := vars["id"]
	ID, err := strconv.ParseInt(doctorId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	doctorDetails, _ := models.GetDoctorsById(ID)
	if (doctorDetails.Morning_count == 0) && (doctorDetails.Afternoon_count == 0) && (doctorDetails.Evening_count == 0) && doctorDetails.Doctor_name == "" {
		message := "Error not found"
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
	} else {
		res, _ := json.Marshal(doctorDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
