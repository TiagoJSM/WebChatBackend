package main

import (
	"fmt"
	"os"

	"github.com/TiagoJSM/WebChatBackend/messages"
	"github.com/labstack/echo/v4"
)

func main() {
	messagesRepository := messages.NewMessagesRepository()
	messagesController := messages.NewMessagesController(messagesRepository)

	e := echo.New()
	e.Static("/static", "assets")
	e.File("/", "public/index.html")
	e.GET("/messages", messagesController.GetAll)
	e.GET("/ws", messagesController.ConnectToSocket)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
