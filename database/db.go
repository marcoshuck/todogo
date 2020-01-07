package database

import (
	"github.com/jinzhu/gorm"
	"github.com/marcoshuck/todogo/errors"
)

type DB struct {
	*gorm.DB
}

func New() (*DB, *errors.Error) {
	db, err := gorm.Open("sqlite3", "todo.db")

	if err != nil {
		return nil, errors.NewError(0, 501, "Cannot open database", &err)
	}

	return &DB{
		DB: db,
	}, nil
}

func Close(db *DB) *errors.Error {
	err := db.Close()
	if err != nil {
		return errors.NewError(0, 501, "Cannot close database", &err)
	}
	return nil
}