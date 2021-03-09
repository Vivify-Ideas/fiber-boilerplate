package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	Password           string `json:"-"`
	ResetPasswordToken string `json:"-"`
}
