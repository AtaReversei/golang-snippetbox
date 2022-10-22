package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errLog:  errLog,
		infoLog: infoLog,
	}

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("starting server on port %s", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
