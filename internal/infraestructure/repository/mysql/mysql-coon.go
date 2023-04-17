package mysql


import (
	"database/sql"
	"log"
)

func RepositoryConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:secret@tcp(mysql-repo:3306)/mysql-db")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}

