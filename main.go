package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func greet(w http.ResponseWriter, r *http.Request) {
	if err := sqlConnection(); err != nil {
		log.Fatalln(err.Error())
	}
	println("SQL Connection working")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func sqlConnection() error {
	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		return err
	}
	defer db.Close()

	path := filepath.Join("db", "init.sql")

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		return ioErr
	}
	sql := string(c)
	if _, execErr := db.Exec(sql); execErr != nil {
		return execErr
	}

	return db.Ping()
}

func main() {
	http.HandleFunc("/greet", greet)
	http.ListenAndServe(":8888", nil)
}
