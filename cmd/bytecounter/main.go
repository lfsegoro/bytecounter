package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bytecounter/internal/service"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := service.New()

	go func() {
		if err := srv.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)

	signal.Notify(sig,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-sig

	cancel()
}
