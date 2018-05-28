package services

import (
	"sample-go-rest/entities"
	"sample-go-rest/repositories"
)

type INewsService interface {
	GetList(params map[string][]string) ([]entities.News, error)
	GetItemById(id int) (*entities.News, error)
	InsertItem(news entities.News) error
	UpdateItem(id int, news entities.News) error
	DeleteItem(id int) error

	AddTopic(newsId int, topicId int) error
	RemoveTopic(newsId int, topicId int) error
}

type newsService struct {
	newsRepository      repositories.INewsRepository
	newsTopicRepository repositories.INewsTopicRepository
}

func NewNewsService(newsRepository repositories.INewsRepository, newsTopicRepository repositories.INewsTopicRepository) INewsService {
	return &newsService{newsRepository: newsRepository, newsTopicRepository: newsTopicRepository}
}

func (service *newsService) GetList(params map[string][]string) ([]entities.News, error) {
	return service.newsRepository.GetList(params)
}

func (service *newsService) GetItemById(id int) (*entities.News, error) {
	return service.newsRepository.GetItemById(id)
}

func (service *newsService) InsertItem(news entities.News) error {
	return service.newsRepository.InsertItem(news)
}

func (service *newsService) UpdateItem(id int, news entities.News) error {
	return service.newsRepository.UpdateItem(id, news)
}

func (service *newsService) DeleteItem(id int) error {
	return service.newsRepository.DeleteItem(id)
}

func (service *newsService) AddTopic(newsId int, topicId int) error {
	return service.newsTopicRepository.InsertRelationship(newsId, topicId)
}

func (service *newsService) RemoveTopic(newsId int, topicId int) error {
	return service.newsTopicRepository.DeleteRelationship(newsId, topicId)
}
