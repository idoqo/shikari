package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gitlab.com/idoko/shikari/api/handler"
	"gitlab.com/idoko/shikari/api/tokens"
	"gitlab.com/idoko/shikari/db"
	"gitlab.com/idoko/shikari/sink"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var (
	apiVersion = "/v1"
	serverAddr = ":8080"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	database := openDbConnection()
	if database.Error != nil {
		logger.Err(database.Error).Msg("could not connect to database")
		os.Exit(1)
	}
	cfg, err := setupKafka(database)
	if err != nil {
		logger.Err(err).Msg("failed to configure kafka clients")
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sink.StartStreaming(ctx, wg, cfg.StreamerConfig)

	wg.Add(1)
	go sink.StartFlushing(ctx, wg, cfg.FlusherConfig)

	jwt := tokens.JWT{}
	h := handler.New(database, logger, jwt)
	r := gin.Default()
	rg := r.Group(apiVersion)
	h.Register(rg)
	r.Run(serverAddr)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	logger.Print(fmt.Sprint(<-sig))
	logger.Print("Stopping Shikari...")
	cancel()
	wg.Wait()
}

func setupKafka(db db.Database) (Config, error){
	var (
		heartbeat = 60 * time.Second
		timeout = 6000
		groupId = "shikari-consumers"
		bootstrapServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	)
	cfg := NewConfig(
		ConfigureFlusher(bootstrapServers, timeout, groupId, db),
		ConfigureStreamer(bootstrapServers, heartbeat),
	)

	return cfg, nil
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