package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	dsn := "root:Shrikant@11@tcp(host.docker.internal:3306)/sys?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	// user := "root"
	// password := "Shrikant@11"
	// hostname := "localhost"
	// port := "3306"
	// dbname := "doctor"
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=Local"
	// d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database Not connected")
		panic(err)
		return
	} else {
		fmt.Println("Database Connected")
	}
	fmt.Printf("Type : %T", d)
	db = d

}
func GetDB() *gorm.DB {
	return db
}
