package messages

type MessagesRepository struct {
	message []message
}

func NewRepository() *MessagesRepository {
	return &MessagesRepository{
		message: make([]message, 0),
	}
}

func (repo *MessagesRepository) GetAll() []message {
	return repo.message
}

func (repo *MessagesRepository) Add(message message) {
	repo.message = append(repo.message, message)
}
