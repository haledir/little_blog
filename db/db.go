package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	migration := `create table if not exists ARTICLES (
        ID integer primary key autoincrement,
        TITLE text not null,
        CONTENT text not null
    );`
	_, err = db.Exec(migration)
	if err != nil {
		return nil, err
	}
	return db, nil
}
