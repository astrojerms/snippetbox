package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

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

	mux := http.NewServeMux()
	// Create a new fileserver  which servces files out of static dir
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server as a handler for all /static/ requests.
	// Strip /static prefix before requests reach the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// The value returned from flag.String() is a pointer to the flag.
	// It must be dereferenced to access the value.
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
