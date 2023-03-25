package database

import (
	"fmt"
	"log"
	"server/entities"

	"github.com/jinzhu/gorm"
)
/**
	\brief type Config for struct in parametr on connection function
*/
type Config struct {
	User     string
	Password string
	DB       string
}
/**
	\brief function GetConnectioString - create a string format "user:password@(localhost:3306)/db.name?..." to connect to DB
*/
var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.DB)
	return connectionString
}

var Connector *gorm.DB

/**
	\brief function Connect - connects with DB
*/
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
		panic("Connect to DB")
	}
	log.Println("Connection was successful!!")
	return nil
}

/**
	\brief function Migrate - function for migration table on DB
*/
func Migrate() {
	Connector.AutoMigrate(entities.Product{})
	Connector.AutoMigrate(entities.Person{})
	log.Println("Database Migration Completed...")
}
