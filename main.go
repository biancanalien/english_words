package main

import (
	"log"
	"net/http"

	"github.com/bianca.pereira/english_words/config"
	"github.com/bianca.pereira/english_words/controller"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := config.Connection()
	defer db.Close()

	vh := &controller.VocabularyHandler{DB: db}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", vh.Add).Methods("POST")
	router.HandleFunc("/", vh.GetAll).Methods("GET")
	router.HandleFunc("/{id}", vh.Update).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
