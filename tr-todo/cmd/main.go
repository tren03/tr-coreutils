package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/tren03/tr-coreutils/tr-todo/boot"
)

func main() {
	app, err := boot.Initalize()
	if err != nil {
		fmt.Println("Error initializing deps", err)
		os.Exit(1)
		return
	}

	port := strconv.Itoa(app.Config.Server.Port)
	go func() {
		if err := app.Server.ListenAndServe(":" + port); err != nil && err != http.ErrServerClosed {
			app.Logger.Error("server failed", "error", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	app.Logger.Info("shutting down gracefully...")
	app.Shutdown(ctx)
}
