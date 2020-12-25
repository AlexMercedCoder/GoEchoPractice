package main

import (
	"gorm.io/gorm"
)

//DECLARE PERSON STRUCT
type Person struct {
	gorm.Model
	Name string `json: name`
	Age int `json: age`
}