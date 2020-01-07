package database

import (
	"github.com/jinzhu/gorm"
	d "github.com/marcoshuck/todogo/definitions"
)

type DB struct {
	*gorm.DB
}

func New() (*DB, *d.Error) {
	db, err := gorm.Open("sqlite3", "todo.db")

	if err != nil {
		return nil, d.NewError(0, 501, "Cannot open database", &err)
	}

	return &DB{
		DB: db,
	}, nil
}

func Close(db *DB) *d.Error {
	err := db.Close()
	if err != nil {
		return d.NewError(0, 501, "Cannot close database", &err)
	}
	return nil
}