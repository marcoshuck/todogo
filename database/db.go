package database

import "github.com/jinzhu/gorm"

type DB struct {
	*gorm.DB
}

func New() *DB {
	db, err := gorm.Open("sqlite3", "todo.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	return &DB{
		DB: db,
	}
}

func Close(db *DB) {
	db.Close()
}