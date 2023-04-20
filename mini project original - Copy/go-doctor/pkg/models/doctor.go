package models

import (
	"github.com/Shrikantgiri25/go-doctor/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Doctor struct {
	gorm.Model

	Doctor_name     string `json:"doctor_name"`
	Morning_count   int    `json:"morning_count"`
	Afternoon_count int    `json:"afternoon_count"`
	Evening_count   int    `json:"evening_count"`
}

type DoctorToBook struct {
	DoctorId  string `json:"doctorId"`
	Batch     string `json:"batch"`
	PatientId string `json:"patientId"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Doctor{})
}

func (d *Doctor) ScheduleDoctor() *Doctor {
	db.NewRecord(d)
	db.Create(&d)
	return d
}

func GetAllDoctors() []Doctor {
	var Doctors []Doctor
	db.Find(&Doctors)
	return Doctors
}
func GetDoctorsById(Id int64) (*Doctor, *gorm.DB) {
	var getDoctor Doctor
	db := db.Where("ID=?", Id).Find(&getDoctor)
	return &getDoctor, db
}

func DeleteDoctor(Id int64) Doctor {
	var doctor Doctor
	db.Where("ID=?", Id).Delete(doctor)
	return doctor
}
