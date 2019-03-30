package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id int
	text string
	done int
)

func main(){
	db, err := sql.Open("mysql",
		"root:Atom2Rock@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, text, done from task")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &text, &done)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, text, done)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}