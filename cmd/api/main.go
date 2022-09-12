package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *log.Logger
}

func openDB(cf config) (*sql.DB, error) {
	// creates an empty connection pool using dsn from config
	db, err := sql.Open("postgres", cf.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cf.db.maxIdleConns)
	db.SetMaxOpenConns(cf.db.maxOpenConns)

	duration, err := time.ParseDuration(cf.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// context with 5-seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// checks if a connection with the database can be made, with a timeout of 5 seconds
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	var conf config

	flag.IntVar(&conf.port, "port", 3333, "Server port")
	flag.StringVar(&conf.env, "env", "development", "Environment (development | staging | production)")
	flag.StringVar(&conf.db.dsn, "db-dsn", "postgres://admin:root@localhost:5432/movie_api?sslmode=disable", "Postgres DSN")
	flag.IntVar(&conf.db.maxIdleConns, "db-max-idle-conns", 25, "Postgres max idle connections")
	flag.IntVar(&conf.db.maxOpenConns, "db-max-open-conns", 25, "Postgres max open connections")
	flag.StringVar(&conf.db.maxIdleTime, "db-max-idle-time", "15m", "Postgres max connection idle time")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(conf)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println("database connected!")

	defer db.Close()

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
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
