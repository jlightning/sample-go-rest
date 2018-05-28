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
}

func NewTopicHandler(newService services.ITopicService) *topicHandler {
	return &topicHandler{topicService: newService}
}

func (handler *topicHandler) Register(mux *mux.Router) {
	mux.HandleFunc("/topics", wrapFunc(handler.handleList)).Methods(http.MethodGet)
	mux.HandleFunc("/topics/{id:[0-9]+}", wrapFunc(handler.handleItem)).Methods(http.MethodGet)
	mux.HandleFunc("/topics", wrapFunc(handler.handleInsert)).Methods(http.MethodPost)
	mux.HandleFunc("/topics/{id:[0-9]+}", wrapFunc(handler.handleUpdate)).Methods(http.MethodPut)
	mux.HandleFunc("/topics/{id:[0-9]+}", wrapFunc(handler.handleDelete)).Methods(http.MethodDelete)
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
