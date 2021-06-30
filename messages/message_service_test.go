package messages

import "testing"

func TestAddingMessage(t *testing.T) {
	messagesRepository := NewRepository()
	messageService := NewService(messagesRepository)
	messageModel := messageModel{"Username", "message"}

	message := messageService.Add(messageModel)

	if message.Username != "Username" {
		t.Errorf("Wrong Username value")
	}
	if message.Text != "message" {
		t.Errorf("Wrong Text value")
	}
}
