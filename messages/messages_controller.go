package messages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MessagesController struct {
	*MessagesRepository
}

type ToDo struct {
	User string
}

func (controller *MessagesController) GetAll(c echo.Context) error {
	messages := controller.MessagesRepository.GetAll()
	return c.Render(http.StatusOK, "homepage.html", messages)
}

func (controller *MessagesController) PostMessage(c echo.Context) error {
	msg := c.FormValue("message")
	controller.MessagesRepository.Add(message{msg})
	return c.JSON(http.StatusOK, "some data")
}