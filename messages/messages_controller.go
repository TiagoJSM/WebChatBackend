package messages

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type MessagesController struct {
	*MessagesService
	clients map[*websocket.Conn]bool
}

func NewController(messageService *MessagesService) *MessagesController {
	return &MessagesController{
		messageService,
		make(map[*websocket.Conn]bool),
	}
}

func (controller *MessagesController) GetAll(c echo.Context) error {
	messages := controller.MessagesService.GetAll()
	return c.JSON(http.StatusOK, messages)
}

func (controller *MessagesController) sendMessageToClients(c echo.Context, messageData *message) {
	for client := range controller.clients {
		err := client.WriteJSON(messageData)
		if err != nil {
			c.Logger().Error(err)
			client.Close()
			delete(controller.clients, client)
		}
	}
}

func (controller *MessagesController) ConnectToSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	controller.clients[ws] = true
	defer ws.Close()

	for {
		msg := messageModel{}

		// Read
		err := ws.ReadJSON(&msg)
		if err != nil {
			c.Logger().Error(err)
			// if user disconnected then stop the message handling
			if closeError := err.(*websocket.CloseError); closeError != nil {
				delete(controller.clients, ws)
				return err
			}
		} else {
			if len(msg.Text) == 0 {
				continue
			}
			messageData := controller.MessagesService.Add(msg)
			// Write
			controller.sendMessageToClients(c, &messageData)
		}
	}
}
