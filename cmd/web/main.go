package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// define cmd line flag 'addr'
	// flag.String() returns a pointer to the flag value
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Create logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server on %s", *addr)
	// we need to dereference addr flag
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
