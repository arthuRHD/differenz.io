package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func sqlConnection(password string, host string) error {
	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf("admin:%s@tcp(%s:3306)/changelogs", password, host),
	)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Ping()
}

func main() {
	if err := sqlConnection("", "changelogs-db.counfmjnq8zy.eu-west-3.rds.amazonaws.com"); err != nil {
		println(err)
	}
	println("SQL Connection working")
	http.HandleFunc("/greet", greet)
	http.ListenAndServe(":8888", nil)
}
