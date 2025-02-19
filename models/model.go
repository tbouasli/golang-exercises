package models

import "gorm.io/gorm"

type People struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}
