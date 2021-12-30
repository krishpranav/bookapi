// DB Setup
package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "sqlite.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Conneted successfully!!")

	database.AutoMigrate(&Book{})

	DB = database
}
