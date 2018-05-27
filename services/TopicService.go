package services

import (
	"sample-go-rest/entities"
	"sample-go-rest/repositories"
)

type ITopicService interface {
	GetList() ([]entities.Topic, error)
	GetItemById(id int) (*entities.Topic, error)
	InsertItem(topic entities.Topic) error
	UpdateItem(id int, topic entities.Topic) error
	DeleteItem(id int) error
	GetListByNewsId(newsId int) ([]entities.Topic, error)
}

type topicService struct {
	topicRepository repositories.ITopicRepository
}

func NewTopicService(topicRepository repositories.ITopicRepository) ITopicService {
	return &topicService{topicRepository: topicRepository}
}

func (service *topicService) GetList() ([]entities.Topic, error) {
	return service.topicRepository.GetList()
}

func (service *topicService) GetItemById(id int) (*entities.Topic, error) {
	return service.topicRepository.GetItemById(id)
}

func (service *topicService) InsertItem(topic entities.Topic) error {
	return service.topicRepository.InsertItem(topic)
}

func (service *topicService) UpdateItem(id int, topic entities.Topic) error {
	return service.topicRepository.UpdateItem(id, topic)
}

func (service *topicService) DeleteItem(id int) error {
	return service.topicRepository.DeleteItem(id)
}

func (service *topicService) GetListByNewsId(newsId int) ([]entities.Topic, error) {
	return service.topicRepository.GetListByNewsId(newsId)
}
