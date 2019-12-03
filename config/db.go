package config

import (
	"fmt"
	"os"

	"github.com/bianca.pereira/english_words/handler"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Connection create new connection with mysql DB
func Connection() *gorm.DB {
	dbURI := getDataBaseURI()
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		fmt.Println("Error while try to connect on mysql ", err)
	}
	db.AutoMigrate(&handler.Vocabulary{})
	return db
}

func getDataBaseURI() string {
	e := godotenv.Load()

	if e != nil {
		fmt.Println("Error while try to load .env file", e)
	}

	username := os.Getenv("db_user")
	pwd := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, pwd, dbHost, dbPort, dbName)
}
