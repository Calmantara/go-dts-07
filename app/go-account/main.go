package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Calmantara/go-account/servers"
	"github.com/Calmantara/go-common/pkg/logger"
)

func main() {
	// run http server
	srv := servers.NewHttpServer()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 2)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	logger.Info(ctx, "Shutdown Server ...")

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(ctx, "Server Shutdown:%v", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
