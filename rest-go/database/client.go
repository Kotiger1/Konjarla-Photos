package database

import (
	"log"
	"rest-go/entity"

	"github.com/jinzhu/gorm"
)

// Connector variable used for CRUD operation's
var Connector *gorm.DB

// Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

// Migrate create/updates database table
func Migrate(table *entity.Picture) {
	//Connector.AutoMigrate(&table)
	Connector.CreateTable(&table)
	log.Println("Table migrated")
}
