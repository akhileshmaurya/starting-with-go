package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func testDbQuery() {
	db, err := sql.Open("mysql", "test:test@/test?charset=utf8")
	checkErr(err)
	rows, err := db.Query("SELECT auth_token FROM user_token")
	checkErr(err)

	for rows.Next() {

		var auth_token string
		err = rows.Scan(&auth_token)
		checkErr(err)
		log.Println("auth_token", auth_token)
	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
