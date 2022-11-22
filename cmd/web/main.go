package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold application wide depdendencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// define 'addr' command line flag with default value :4000
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Use flag.parse to parse command line flag. need to call before using
	// addr variable.
	flag.Parse()

	// Create new logger for informational logging. 3 arguments:
	// path to wrtie logs, prefix, additional information to include
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create error log like above, including shortFile (filename/number)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// init a new instance of the application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// The value returned from flag.String() is a pointer to the flag.
	// It must be dereferenced to access the value.
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
