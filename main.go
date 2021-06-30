package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/TiagoJSM/WebChatBackend/messages"
	"github.com/labstack/echo/v4"
)

func main() {
	messagesRepository := messages.NewRepository()
	messagesController := messages.NewController(messagesRepository)

	e := echo.New()
	e.Static("/static", "assets")
	e.File("/", "public/index.html")
	e.GET("/messages", messagesController.GetAll)
	e.GET("/ws", messagesController.ConnectToSocket)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
