package main

import (
	"context"
	"go_logs/internal/file"
	"go_logs/internal/udpserver"
	"go_logs/pkg/config"
	"go_logs/pkg/logger"
	"log"
	"os"
	"os/signal"
	"time"
)

func run() error {
	_ = config.Get()
	l := logger.Get("logs/log.log")
	l.Info().Msgf("logger init %s", "ok")

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		//l.Info().Msg("Stopping...")
		cancel()
		time.Sleep(50 * time.Millisecond)
		// time.Sleep(2 * time.Second)
		log.Printf("Stopped.")
	}()

	log.Printf("ctx %+v", ctx)

	f := file.New(ctx, l)
	udpserver.New(ctx, l, f)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatalf("err: %s", err.Error())
	}
}
