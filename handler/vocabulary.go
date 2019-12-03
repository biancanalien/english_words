package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type VocabularyHandler struct {
	DB *gorm.DB
}

//Vocabulary represents a model of vocabulary in the system
type Vocabulary struct {
	gorm.Model
	Expression  string `gorm:"unique;not null"`
	Meaning     string
	Translation string
}

//Add create new vocabulary
func (vh *VocabularyHandler) Add(w http.ResponseWriter, r *http.Request) {
	var v Vocabulary
	_ = json.NewDecoder(r.Body).Decode(&v)
	res := vh.DB.Create(&v)

	if res.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
	} else {
		respondWithJSON(w, http.StatusOK, v)
	}
}

//Update update some vocabulary by id
func (vh *VocabularyHandler) Update(w http.ResponseWriter, r *http.Request) {
	var new Vocabulary
	json.NewDecoder(r.Body).Decode(&new)

	var v Vocabulary
	params := mux.Vars(r)
	vh.DB.First(&v, params["id"])

	if res := vh.DB.Model(&v).Update(new); res.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
	} else {
		respondWithJSON(w, http.StatusOK, v)
	}
}

//GetAll returns all vocabularies
func (vh *VocabularyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var vs []Vocabulary
	if res := vh.DB.Find(&vs); res.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error")
	} else {
		respondWithJSON(w, http.StatusOK, vs)
	}
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
