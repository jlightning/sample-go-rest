package handlers

import (
	"net/http"
	"sample-go-rest/services"
	"encoding/json"
	"io/ioutil"
	"sample-go-rest/entities"
	"github.com/gorilla/mux"
)

type topicHandler struct {
	topicService services.ITopicService
	newsService  services.INewsService
}

func NewTopicHandler(newService services.ITopicService, newsService services.INewsService) *topicHandler {
	return &topicHandler{topicService: newService, newsService: newsService}
}

func (handler *topicHandler) Register(mux *mux.Router) {
	mux.HandleFunc("/topics", wrapFunc(handler.handleList)).Methods(http.MethodGet)
	mux.HandleFunc("/topics/{id:[0-9]+}", wrapFunc(handler.handleItem)).Methods(http.MethodGet)
	mux.HandleFunc("/topics", wrapFunc(handler.handleInsert)).Methods(http.MethodPost)
	mux.HandleFunc("/topics/{id:[0-9]+}", wrapFunc(handler.handleUpdate)).Methods(http.MethodPut)
	mux.HandleFunc("/topics/{id:[0-9]+}", wrapFunc(handler.handleDelete)).Methods(http.MethodDelete)

	mux.HandleFunc("/topics/{id:[0-9]+}/news", wrapFunc(handler.handleListNews)).Methods(http.MethodGet)
	mux.HandleFunc("/topics/{id:[0-9]+}/news/{newsId:[0-9]+}", wrapFunc(handler.handleItemNews)).Methods(http.MethodGet)
	mux.HandleFunc("/topics/{id:[0-9]+}/news/{newsId:[0-9]+}", wrapFunc(handler.handleAddTopicRelationship)).Methods(http.MethodPost)
	mux.HandleFunc("/topics/{id:[0-9]+}/news/{newsId:[0-9]+}", wrapFunc(handler.handleDeleteTopicRelationship)).Methods(http.MethodDelete)
}

func (handler *topicHandler) handleList(request *http.Request) (interface{}, error) {
	return handler.topicService.GetList(request.URL.Query())
}

func (handler *topicHandler) handleItem(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}
	return handler.topicService.GetItemById(id)
}

func (handler *topicHandler) handleInsert(request *http.Request) (interface{}, error) {
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var topic entities.Topic
	err = json.Unmarshal(requestBody, &topic)
	if err != nil {
		return nil, err
	}

	return nil, handler.topicService.InsertItem(topic)
}

func (handler *topicHandler) handleUpdate(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var topic entities.Topic
	err = json.Unmarshal(requestBody, &topic)
	if err != nil {
		return nil, err
	}

	return nil, handler.topicService.UpdateItem(id, topic)
}

func (handler *topicHandler) handleDelete(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	return nil, handler.topicService.DeleteItem(id)
}

func (handler *topicHandler) handleAddTopicRelationship(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	newsId, err := getRoueParamFromRequest("newsId", request)
	if err != nil {
		return nil, err
	}

	return nil, handler.newsService.AddTopic(newsId, id)
}

func (handler *topicHandler) handleDeleteTopicRelationship(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	newsId, err := getRoueParamFromRequest("newsId", request)
	if err != nil {
		return nil, err
	}

	return nil, handler.newsService.RemoveTopic(newsId, id)
}

func (handler *topicHandler) handleListNews(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	return handler.newsService.GetListByTopicId(id, request.URL.Query())
}

func (handler *topicHandler) handleItemNews(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	newsId, err := getRoueParamFromRequest("newsId", request)
	if err != nil {
		return nil, err
	}

	return handler.newsService.GetItemByTopicId(id, newsId)
}
