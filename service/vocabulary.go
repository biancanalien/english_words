package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Vocabulary represents a model of vocabulary in the system
type Vocabulary struct {
	gorm.Model
	Expression  string `gorm:"unique;not null"`
	Meaning     string
	Translation string
}

func AddVocabulary(w http.ResponseWriter, r *http.Request) {
	var v models.Vocabulary
	_ = json.NewDecoder(r.Body).Decode(&v)
	res := db.Create(&v)

	if res.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
	} else {
		respondWithJSON(w, http.StatusOK, v)
	}
}

func UpdateVocabulary(w http.ResponseWriter, r *http.Request) {
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

func GetAllVocabularies(w http.ResponseWriter, r *http.Request) {
	var vs []models.Vocabulary
	if res := db.Find(&vs); res.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
	} else {
		respondWithJSON(w, http.StatusOK, vs)
	}
}
