package database

import "github.com/jinzhu/gorm"

// Model defines a basic struct for new database entries
type Model struct {
	gorm.Model
}