package services

import (
	"sample-go-rest/entities"
	"sample-go-rest/repositories"
)

type ITopicService interface {
	GetList(params map[string][]string) ([]entities.Topic, error)
	GetItemById(id int) (*entities.Topic, error)
	InsertItem(topic entities.Topic) error
	UpdateItem(id int, topic entities.Topic) error
	DeleteItem(id int) error
	GetListByNewsId(newsId int, params map[string][]string) ([]entities.Topic, error)
	GetItemByNewsId(newsId int, topicId int) (*entities.Topic, error)
}

type topicService struct {
	topicRepository repositories.ITopicRepository
}

func NewTopicService(topicRepository repositories.ITopicRepository) ITopicService {
	return &topicService{topicRepository: topicRepository}
}

func (service *topicService) GetList(params map[string][]string) ([]entities.Topic, error) {
	return service.topicRepository.GetList(params)
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

func (service *topicService) GetListByNewsId(newsId int, params map[string][]string) ([]entities.Topic, error) {
	return service.topicRepository.GetListByNewsId(newsId, params)
}

func (service *topicService) GetItemByNewsId(newsId int, topicId int) (*entities.Topic, error) {
	return service.topicRepository.GetItemByNewsId(newsId, topicId)
}
