package config

import (
	"fmt"

	"github.com/bianca.pereira/english_words/models"
	"github.com/jinzhu/gorm"
)

// Connection create new connection with mysql DB
func Connection() *gorm.DB {
	db, err := gorm.Open("mysql", "")
	if err != nil {
		fmt.Println("Error while try to connect on mysql ", err)
	}
	db.AutoMigrate(&models.Vocabulary{})
	return db
}
