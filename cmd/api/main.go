package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var conf config

	flag.IntVar(&conf.port, "port", 3333, "Server port")
	flag.StringVar(&conf.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: conf,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on port %d...", conf.env, conf.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
