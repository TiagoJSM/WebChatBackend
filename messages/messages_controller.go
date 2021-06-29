package messages

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type MessagesController struct {
	*MessagesRepository
	clients map[*websocket.Conn]bool
}

func NewMessagesController(messageRepository *MessagesRepository) *MessagesController {
	return &MessagesController{
		messageRepository,
		make(map[*websocket.Conn]bool),
	}
}

func (controller *MessagesController) GetAll(c echo.Context) error {
	messages := controller.MessagesRepository.GetAll()
	return c.JSON(http.StatusOK, messages)
}

func (controller *MessagesController) PostMessage(c echo.Context) error {
	msg := c.FormValue("message")

	if msg == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid message content"}
	}

	controller.MessagesRepository.Add(message{"Username", time.Now(), msg})
	return c.JSON(http.StatusCreated, msg)
}

func (controller *MessagesController) ConnectToSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	controller.clients[ws] = true
	defer ws.Close()

	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}

		messageData := message{"Username", time.Now(), string(msg)}
		controller.MessagesRepository.Add(messageData)
		// Write
		for client := range controller.clients {
			err = client.WriteJSON(messageData)
			if err != nil {
				c.Logger().Error(err)
				client.Close()
				delete(controller.clients, client)
			}
		}
	}
}
