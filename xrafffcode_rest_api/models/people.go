package models

import "gorm.io/gorm"

type People struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	University string `json:"university"`
	Hobbies    string `json:"hobbies"`
	Portfolio  string `json:"portfolio"`
}
