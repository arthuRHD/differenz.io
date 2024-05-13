package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/arthuRHD/differenz.io/api"
	"github.com/arthuRHD/differenz.io/db"
)

func greet(w http.ResponseWriter, r *http.Request) error {
	if err := sqlConnection(); err != nil {
		return err
	}
	println("SQL Connection working")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
	return nil
}

func sqlConnection() error {
	conn, err := db.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := db.GenerateSchema(conn); err != nil {
		return err
	}

	return conn.Ping()
}

func main() {
	http.HandleFunc("/greet", api.ErrorMiddleware(greet))
	http.ListenAndServe(":8888", nil)
}
