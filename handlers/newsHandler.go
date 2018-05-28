package handlers

import (
	"net/http"
	"sample-go-rest/services"
	"encoding/json"
	"strconv"
	"io/ioutil"
	"sample-go-rest/entities"
	"github.com/gorilla/mux"
)

type newsHandler struct {
	newsService  services.INewsService
	topicService services.ITopicService
}

func NewNewsHandler(newService services.INewsService, topicService services.ITopicService) *newsHandler {
	return &newsHandler{newsService: newService, topicService: topicService}
}

func (handler *newsHandler) Register(mux *mux.Router) {
	mux.HandleFunc("/news", wrapFunc(handler.handleList)).Methods(http.MethodGet)
	mux.HandleFunc("/news/{id:[0-9]+}", wrapFunc(handler.handleItem)).Methods(http.MethodGet)
	mux.HandleFunc("/news", wrapFunc(handler.handleInsert)).Methods(http.MethodPost)
	mux.HandleFunc("/news/{id:[0-9]+}", wrapFunc(handler.handleUpdate)).Methods(http.MethodPut)
	mux.HandleFunc("/news/{id:[0-9]+}", wrapFunc(handler.handleDelete)).Methods(http.MethodDelete)

	mux.HandleFunc("/news/{id:[0-9]+}/topics", wrapFunc(handler.handleListTopic)).Methods(http.MethodGet)
	mux.HandleFunc("/news/{id:[0-9]+}/topics/{topicId:[0-9]+}", wrapFunc(handler.handleItemTopic)).Methods(http.MethodGet)
	mux.HandleFunc("/news/{id:[0-9]+}/topics/{topicId:[0-9]+}", wrapFunc(handler.handleAddTopicRelationship)).Methods(http.MethodPut)
	mux.HandleFunc("/news/{id:[0-9]+}/topics/{topicId:[0-9]+}", wrapFunc(handler.handleDeleteTopicRelationship)).Methods(http.MethodDelete)
}

func (handler *newsHandler) handleList(request *http.Request) (interface{}, error) {
	params := request.URL.Query()
	return handler.newsService.GetList(params)
}

func (handler *newsHandler) handleItem(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}
	return handler.newsService.GetItemById(id)
}

func (handler *newsHandler) handleInsert(request *http.Request) (interface{}, error) {
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var news entities.News
	err = json.Unmarshal(requestBody, &news)
	if err != nil {
		return nil, err
	}

	return nil, handler.newsService.InsertItem(news)
}

func (handler *newsHandler) handleUpdate(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var news entities.News
	err = json.Unmarshal(requestBody, &news)
	if err != nil {
		return nil, err
	}

	return nil, handler.newsService.UpdateItem(id, news)
}

func (handler *newsHandler) handleDelete(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	return nil, handler.newsService.DeleteItem(id)
}

func (handler *newsHandler) handleAddTopicRelationship(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	topicId, err := getRoueParamFromRequest("topicId", request)
	if err != nil {
		return nil, err
	}

	return nil, handler.newsService.AddTopic(id, topicId)
}

func (handler *newsHandler) handleDeleteTopicRelationship(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	topicId, err := getRoueParamFromRequest("topicId", request)
	if err != nil {
		return nil, err
	}

	return nil, handler.newsService.RemoveTopic(id, topicId)
}

func (handler *newsHandler) handleListTopic(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	return handler.topicService.GetListByNewsId(id)
}

func (handler *newsHandler) handleItemTopic(request *http.Request) (interface{}, error) {
	id, err := getRoueParamFromRequest("id", request)
	if err != nil {
		return nil, err
	}

	topicId, err := getRoueParamFromRequest("topicId", request)
	if err != nil {
		return nil, err
	}

	return handler.topicService.GetItemByNewsId(id, topicId)
}

func getRoueParamFromRequest(key string, request *http.Request) (int, error) {
	vars := mux.Vars(request)
	if idStr, ok := vars[key]; ok {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, newHttpError("bad request", 400)
}
