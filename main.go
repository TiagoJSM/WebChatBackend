package main

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/TiagoJSM/WebChatBackend/messages"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	messagesRepository := messages.NewMessagesRepository()
	messagesController := messages.MessagesController{messagesRepository}

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return messagesController.GetAll(c)
	})
	e.POST("/message", func(c echo.Context) error {
		return messagesController.PostMessage(c)
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
