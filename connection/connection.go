package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/lib/pq"
)

/// Подключение к базе Постгри

//const (
//	host     = "host"
//	user     = "user"
//	database = "database"
//	port     = "port"
//	password = "password"
//)
//
//func OpenConnection() *sql.DB {
//	psglInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user,
//		password, database)
//	db, err := sql.Open("postgres", psglInfo)
//	if err != nil {
//		fmt.Print("1", err)
//	}
//
//	err = db.Ping()
//	if err != nil {
//		fmt.Print("2", err)
//	}
//
//	return db
//}

///---------------------------------Подключение SQLite3----------------------------------

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "database/Task_DB")
	if err != nil {
		fmt.Println("1", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("2", err)
	}
	return db
}
