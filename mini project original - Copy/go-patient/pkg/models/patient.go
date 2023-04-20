package models

import (
	"strconv"

	"github.com/Shrikantgiri25/go-patient/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Patient struct {
	gorm.Model

	Patient_name string `json:"patient_name"`
	Age          int    `json:"age"`
	Batch        string `json:"batch"`
	DoctorId     string `json:"doctorId"`
	Doctor_name  string `json:"doctor_name"`
	Booked       bool   `json: "booked"`
}

type DoctorToBook struct {
	PatientId string `json:"patientId"`
	DoctorId  string `json:"doctorId"`
	Batch     string `json:"batch"`
}
type Doctor struct {
	gorm.Model

	Doctor_name     string `json:"doctor_name"`
	Morning_count   int    `json:"morning_count"`
	Afternoon_count int    `json:"afternoon_count"`
	Evening_count   int    `json:"evening_count"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Patient{})
}

func (p *Patient) AddPatient() *Patient {
	db.NewRecord(p)
	db.Create(&p)
	return p
}
func GetAllPatients() []Patient {
	var Patients []Patient

	// query := db.Table("doctors").Select("doctors.doctor_name").Joins("left join patients on doctors.id = patients.doctor_id").Group("doctors.id")
	// db.Model(&Patient{}).Joins("join (?) q on doctors.doctor_name = q.doctor_name", query).Scan(&Patients)
	//var Doctors []Doctor
	// db.Table("doctors").Select("doctors.doctor_name").Joins("left join patients on patients.doctor_id = doctors.id").Scan(&Patients)
	// db.Find(&Patients)
	//db.Model(&Patient{}).Select("*").Joins("JOIN doctors ON patients.doctor_id = doctors.id or patients.doctor_Name = doctors.doctor_name").Find(&Patients)
	//db.Preload("doctors").Find(&Patients)
	//db.Table("doctors").Joins("join patients on patients.doctor_id = doctors.id").Scan(&Patients)
	//var Doctors []Doctor
	//.InnerJoins("doctors").Find(&Patients)
	//db.Table("doctors").Select("doctors.doctor_name, doctors.id").Joins("left join patients on patients.doctor_id = doctors.id and patients.doctor_name = doctors.doctor_name").Scan(&Doctors)
	db.Table("doctors").Select("patients.id, patients.patient_name, patients.age, patients.booked ,patients.batch, patients.doctor_id, doctors.doctor_name").Joins("inner join patients on patients.doctor_id = doctors.id ").Scan(&Patients)

	//This part is working fine ???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
	// db.Find(&Patients)
	// for i := 0; i < len(Patients); i++ {
	// 	getId, _ := strconv.ParseUint(Patients[i].DoctorId, 0, 0)
	// 	var docname Doctor
	// 	db.Table("doctors").Where("doctors.id=?", getId).Joins("left join patients on patients.doctor_id = doctors.id").Scan(&docname)
	// 	Patients[i].DoctorName = docname.Doctor_name
	// }
	//?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????

	//db.Joins("JOIN doctors ON Patients.doctor_id = doctors.id").Find(&Patients)
	// for _, pval := range Patients {
	// 	for _, dval := range Doctors {
	// 		if pval.DoctorId == dval. {
	// 			pval.DoctorName = dval.Doctor_name
	// 		}
	// 	}
	// }
	// //return Doctors
	//db.Table("patients").Select("patients.*, doctors.doctor_name as patients.doctor_name").Joins("left join doctors on doctors.id = patients.doctor_id").Scan(&Patients)
	//db.Joins("doctors").Find(&Patients)
	return Patients

}

func GetPatientById(Id int64) (*Patient, *gorm.DB) {
	var getDoctor Patient
	var docname Doctor
	kb := db.Where("ID=?", Id).Find(&getDoctor)
	getId, _ := strconv.ParseUint(getDoctor.DoctorId, 0, 0)
	db.Table("doctors").Where("doctors.id=?", getId).Joins("left join patients on patients.doctor_id = doctors.id").Scan(&docname)
	getDoctor.Doctor_name = docname.Doctor_name
	return &getDoctor, kb

	//db.Where("doctors.id = ?", getDoctor.DoctorId).Joins("JOIN doctors ON patients.doctor_name = doctors.doctor_name").Find(&docname)
	//getDoctor.DoctorName = docname.DoctorName
	// return &getDoctor, kb
	//db.Joins("JOIN doctors ON patients.doctor_id = doctors.id").Find(&docname)
	//db.Joins("JOIN doctors ON doctors.id = patients.doctor_id").Find(&docname)
	//db.Table("doctors").Where("doctors.id=?", getId).Joins("left join patients on patients.doctor_id = doctors.id and patients.doctor_name = doctors.doctor_name").Scan(&docname) (worked)
}

func DeletePatient(Id int64) Patient {
	var patient Patient
	db.Where("ID=?", Id).Delete(patient)
	return patient
}

func GetDocBatch(pdocId string, pbatch string) []Patient {
	var Patients []Patient
	if pbatch == "" {
		db.Where("doctor_id = ?", pdocId).Find(&Patients)
		return Patients
	} else if pdocId == "" {
		db.Where("Batch = ?", pbatch).Find(&Patients)
		//db.Where("doctor_id = ?", pdocId).Find(&Patients)
		return Patients
	} else if pbatch == "" && pdocId == "" {
		db.Find(&Patients)
		return Patients
	} else {
		db.Where("doctor_id = ? AND batch = ?", pdocId, pbatch).Find(&Patients)
		return Patients
	}
	//db.Where("Batch = ?", pbatch).Find(&Patients)
}

// func GetJoinDoc(pdocId string) string {
// 	//db.Table("doctors").Select("doctor_name as latest").Joins("left join patient patient on patient.id = patient.id").Group("doctors.doctor_name")
// 	var doc Doctor
// 	db.Joins("doctors").Find(&Patient{})
// 	docId, _ := strconv.ParseInt(pdocId, 0, 0)
// 	db.Where("doctor_id = ? ", docId).Find(&doc)
// 	return doc.Doctor_name
// }

// func GetSameDoc(pdocId string) []Patient {
// 	var Patients []Patient
// 	//db.Where("Batch = ?", pbatch).Find(&Patients)
// 	db.Where("doctor_id = ?", pdocId).Find(&Patients)
// 	return Patients
// }

// func GetSameBatch(pbatch string) []Patient {
// 	var Patients []Patient
// 	db.Where("Batch = ?", pbatch).Find(&Patients)
// 	//db.Where("doctor_id = ?", pdocId).Find(&Patients)
// 	return Patients
// }
