package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tren03/tr-coreutils/tr-todo/boot"
)

func main() {
	app, err := boot.Initalize()
	if err != nil {
		fmt.Println("Error initializing deps", err)
	}
	//fmt.Printf("%#v", app.Repos.Todo.CreateTodo(&todo.Todo{}))
	// go startServer(app)

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan // Block until signal received

	app.Logger.Info("shutting down gracefully...")
	app.Shutdown()
}
