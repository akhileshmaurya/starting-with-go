package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "komli:komli@/PredictModel?charset=utf8")
	checkErr(err)
	rows, err := db.Query("SELECT pj_cron_expression FROM predictmodel_jobs")
	checkErr(err)

	for rows.Next() {

		var pj_cron_expression string
		err = rows.Scan(&pj_cron_expression)
		checkErr(err)
		fmt.Println("pj_cron_expression", pj_cron_expression)
	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
