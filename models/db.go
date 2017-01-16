package models

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type DbConnector struct {
    path string
}

func NewConnector(value string) DbConnector {
    return DbConnector{path: value}
}

func (d *DbConnector) initDB() *sql.DB {
    db, err := sql.Open("sqlite3", d.path)

    if err != nil {
        panic(err)
    }
    if db == nil {
        panic("db nil")
    }

    // create tables
    d.createTables(db)

    return db
}

func (d *DbConnector) createTables(db *sql.DB) {
    // create table if not exists
    sqlTable := `
	CREATE TABLE IF NOT EXISTS files(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		hash VARCHAR(64) NULL,
        name VARCHAR(64) NULL,
        path VARCHAR(64) NULL,
        type VARCHAR(64) NULL,
        description TEXT NULL,
        awss3 BOOLEAN NOT NULL DEFAULT 0
	);
	`

    _, err := db.Exec(sqlTable)
    if err != nil {
        panic(err)
    }
}
