package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Shrikantgiri25/go-patient/pkg/models"
	"github.com/Shrikantgiri25/go-patient/pkg/utils"

	"github.com/gorilla/mux"
)

var NewPatient models.Patient

func GetPatients(w http.ResponseWriter, r *http.Request) {
	newPatients := models.GetAllPatients()
	res, _ := json.Marshal(newPatients)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func AddPatient(w http.ResponseWriter, r *http.Request) {
	addPatient := &models.Patient{}
	utils.ParseBody(r, addPatient)
	if addPatient.Patient_name == "" || addPatient.Age == 0 {
		message := "Patient Name or Age can't be Null"
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	} else {
		b := addPatient.AddPatient()

		res, _ := json.Marshal(b)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetPatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientId := vars["id"]
	ID, err := strconv.ParseInt(patientId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	patientDetails, _ := models.GetPatientById(ID)
	if (patientDetails.Patient_name == "") && (patientDetails.Age == 0) && (patientDetails.Batch == "") && (patientDetails.DoctorId == "") {
		message := "Error not found"
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)

	} else {
		res, _ := json.Marshal(patientDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientId := vars["id"]
	ID, err := strconv.ParseInt(patientId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	patientDetails, _ := models.GetPatientById(ID)
	if patientDetails.Booked == true {
		removeDoctorToBook := &models.DoctorToBook{}
		removeDoctorToBook.DoctorId = patientDetails.DoctorId
		removeDoctorToBook.Batch = patientDetails.Batch
		js, err := json.Marshal(removeDoctorToBook)
		if err != nil {
			fmt.Println("Cant convert to JSON")
		}
		responseBody := bytes.NewBuffer(js)

		resp, err := http.Post("http://doctors:2010/cancelSlot", "application/json", responseBody)
		if err != nil {
			fmt.Println("Error while sending request to doctor book slot", err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		_ = models.DeletePatient(ID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	} else if (patientDetails.Patient_name == "") && (patientDetails.Age == 0) && (patientDetails.Batch == "") && (patientDetails.DoctorId == "") {
		message := "Error not found"
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
	} else {
		_ = models.DeletePatient(ID)
		message := "Deleted"
		res, _ := json.Marshal(message)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func BookSchedule(w http.ResponseWriter, r *http.Request) {
	doctorToBook := &models.DoctorToBook{}
	utils.ParseBody(r, doctorToBook)

	js, err := json.Marshal(doctorToBook)
	if err != nil {
		fmt.Println("Cant convert to json", err)
	}
	patientID := doctorToBook.PatientId
	ID, err := strconv.ParseInt(patientID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	patientDetails, db := models.GetPatientById(ID)
	if patientDetails.Booked == true {
		message := "Patient have already booked the batch."
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(res)
		return
	}

	responseBody := bytes.NewBuffer(js)
	resp, err := http.Post("http://doctors:2010/bookSlot", "application/json", responseBody)

	if err != nil {
		fmt.Println("Error while Sending request to Teacher Book Slot", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	if strBody != "\"Cannot Book. It is already Full\"" {

		patientDetails.Batch = doctorToBook.Batch
		patientDetails.DoctorId = doctorToBook.DoctorId
		patientDetails.Booked = true
		db.Save(&patientDetails)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

	defer resp.Body.Close()
}

// ////////////////////////////////////////////////////////////////////////// New functions
func BookSlot(w http.ResponseWriter, r *http.Request) {
	doctorToBook := &models.DoctorToBook{}
	vars := mux.Vars(r)
	patientId := vars["pid"]
	doctorId := vars["did"]
	batch := vars["batch"]
	doctorToBook.PatientId = patientId
	doctorToBook.DoctorId = doctorId
	doctorToBook.Batch = batch
	utils.ParseBody(r, doctorToBook)

	js, err := json.Marshal(doctorToBook)
	if err != nil {
		fmt.Println("Cant convert to json", err)
	}
	PID, err := strconv.ParseInt(patientId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	patientDetails, db := models.GetPatientById(PID)

	if patientDetails.Booked == true {
		message := "Patient have already booked the batch."
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(res)
		return
	}

	responseBody := bytes.NewBuffer(js)
	resp, err := http.Post("http://doctors:2010/bookSlot", "application/json", responseBody)

	if err != nil {
		fmt.Println("Error while Sending request to doctor Book Slot", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	if strBody != "\"Cannot Book. It is already Full\"" {

		patientDetails.Batch = doctorToBook.Batch
		patientDetails.DoctorId = doctorToBook.DoctorId
		// site := "http://doctors:2010/doctorName/" + doctorToBook.DoctorId
		// resp, err := http.Get(site)
		// if err != nil {
		// 	message := "Error while sending request to doctor module"
		// 	res, _ := json.Marshal(message)
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusNotAcceptable)
		// 	w.Write(res)
		// 	return
		// }
		// addDoc := &models.Doctor{}
		// resData, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	message := "Error while Unmarshalling"
		// 	res, _ := json.Marshal(message)
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusNotAcceptable)
		// 	w.Write(res)
		// 	return
		// }
		// json.Unmarshal(resData, &addDoc)
		// patientDetails.DoctorName = addDoc.Doctor_name

		// //res, _ := json.Marshal(resp.Body)
		// // body, _ := ioutil.ReadAll(resp.Body)
		// // strBody := string(body)
		// // var docname rune
		// // for _, val := range strBody {
		// // 	if val == '"' {
		// // 		docname = docname + val
		// // 	}
		// // }
		// var docId string
		// docId = doctorToBook.DoctorId
		// patientDetails.DoctorName = models.getJoinDoc(docId)
		patientDetails.Booked = true
		db.Save(&patientDetails)
	} else {
		message := "Batch is full."
		res, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

	defer resp.Body.Close()
}

func GetSameDocBatch(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	DocId := params.Get("DoctorId")
	batch := params.Get("Batch")
	if DocId == "" && batch == "" {
		newPatients := models.GetAllPatients()
		res, _ := json.Marshal(newPatients)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		newPatients := models.GetDocBatch(DocId, batch)
		res, _ := json.Marshal(newPatients)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// func GetSameDoc(w http.ResponseWriter, r *http.Request) {
// 	params := r.URL.Query()
// 	DocId := params.Get("DoctorId")
// 	//batch := params.Get("Batch")

// 	newPatients := models.GetSameDoc(DocId)
// 	res, _ := json.Marshal(newPatients)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// //getting same batch

// func GetSameBatch(w http.ResponseWriter, r *http.Request) {
// 	params := r.URL.Query()
// 	//DocId := params.Get("DoctorId")
// 	batch := params.Get("Batch")

// 	newPatients := models.GetSameBatch(batch)
// 	res, _ := json.Marshal(newPatients)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
