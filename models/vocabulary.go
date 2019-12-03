package models

import (
	"github.com/jinzhu/gorm"
)

//Vocabulary model
type Vocabulary struct {
	gorm.Model
	Expression  string `gorm:"unique;not null"`
	Meaning     string
	Translation string
}
