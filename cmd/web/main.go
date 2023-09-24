package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"uwDavid/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHost  = "localhost"
	DBPort  = "3306"
	DBUser  = "web"
	DBPass  = "pass"
	DBDbase = "snippetbox"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// define cmd line flag 'addr'
	// flag.String() returns a pointer to the flag value
	addr := flag.String("addr", ":4000", "HTTP network address")
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBDbase)
	log.Println(dbConn)
	dsn := flag.String("dsn", dbConn, "MySQL database")
	flag.Parse()

	// Create logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server on %s", *addr)
	// we need to dereference addr flag
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
