package models

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

const path = "./storage/goup.db"

func initDB() *sql.DB {
    db, err := sql.Open("sqlite3", path)

    if err != nil {
        panic(err)
    }
    if db == nil {
        panic("db nil")
    }

    // create tables
    createTables(db)

    return db
}

func createTables(db *sql.DB) {
    // create table if not exists
    sqlTable := `
	CREATE TABLE IF NOT EXISTS files(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		hash VARCHAR(64) NULL,
        name VARCHAR(64) NULL,
        path VARCHAR(64) NULL,
        type VARCHAR(64) NULL,
        description TEXT NULL
	);
	`

    _, err := db.Exec(sqlTable)
    if err != nil {
        panic(err)
    }
}
