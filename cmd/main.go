package main

import (
	"context"
	"fmt"
	"gitlab.com/idoko/shikari/db"
	"gitlab.com/idoko/shikari/sink"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sink.StartStreaming(ctx, wg)

	database := openDbConnection()
	if database.Error != nil {
		log.Fatal(database.Error)
		return
	}
	
	wg.Add(1)
	go sink.StartFlushing(ctx, wg, database)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	log.Println(fmt.Sprint(<-sig))
	log.Println("Stopping Shikari...")

	cancel()
	wg.Wait()
}

func openDbConnection() db.Database {
	var d db.Database
	dbHost, port, dbUser, dbPass, dbName, dbSSLMode :=
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSLMODE")
	dbPort, err := strconv.Atoi(port)
	if err != nil {
		d.Error = fmt.Errorf("Port is invalid: %s", err.Error())
	} else {
		d = db.Connect(dbHost, dbUser, dbPass, dbName, dbSSLMode, dbPort)
	}
	return d
}