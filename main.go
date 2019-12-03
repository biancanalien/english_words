package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bianca.pereira/english_words/config"

	"github.com/bianca.pereira/english_words/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func addVocabulary(w http.ResponseWriter, r *http.Request) {
	var v models.Vocabulary
	_ = json.NewDecoder(r.Body).Decode(&v)
	res := db.Create(&v)

	if res.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
	} else {
		respondWithJSON(w, http.StatusOK, v)
	}
}

func updateVocabulary(w http.ResponseWriter, r *http.Request) {
	var new models.Vocabulary
	json.NewDecoder(r.Body).Decode(&new)

	var v models.Vocabulary
	params := mux.Vars(r)
	db.First(&v, params["id"])

	if res := db.Model(&v).Update(new); res.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
	} else {
		respondWithJSON(w, http.StatusOK, v)
	}
}

func getVocabulary(w http.ResponseWriter, r *http.Request) {
	var vs []models.Vocabulary
	if res := db.Find(&vs); res.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
	} else {
		respondWithJSON(w, http.StatusOK, vs)
	}
}

func main() {
	db = config.Connection()
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/create", addVocabulary).Methods("POST")
	router.HandleFunc("/getAll", getVocabulary).Methods("GET")
	router.HandleFunc("/update/{id}", updateVocabulary).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
