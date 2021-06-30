package messages

import "time"

type MessagesService struct {
	*MessagesRepository
}

func NewService(messageRepository *MessagesRepository) *MessagesService {
	return &MessagesService{
		messageRepository,
	}
}

func (service *MessagesService) GetAll() []message {
	return service.MessagesRepository.GetAll()
}

func (service *MessagesService) Add(messageModel messageModel) message {
	messageData := message{messageModel.Username, time.Now(), messageModel.Text}
	service.MessagesRepository.Add(messageData)
	return messageData
}
