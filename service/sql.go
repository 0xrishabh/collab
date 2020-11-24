package service

import (
	_ "fmt"
	_ "strconv"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

func Save(RequestType string, Ip string, Headers string){

	database, _ := sql.Open("sqlite3", "./logs.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY, RequestType TEXT, Ip TEXT, Headers TEXT)")
    statement.Exec()
	statement, _ = database.Prepare("INSERT INTO logs (RequestType, Ip, Headers) VALUES (?, ?, ?)")
    statement.Exec(RequestType,Ip,Headers)
	/*
		rows, _ := database.Query("SELECT id, RequestType, Ip, headers FROM people")
	    var id int
	    var firstname string
	    var lastname string
	    for rows.Next() {
	        rows.Scan(&id, &firstname, &lastname)
	        fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	    }
    */
}