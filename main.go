package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"sample-go-rest/handlers"
	"database/sql"
	"sample-go-rest/repositories"
	"sample-go-rest/services"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/sample_rest?charset=utf8")
	if err != nil {
		panic(err)
	}
	mux := mux.NewRouter().StrictSlash(true)

	newsRepository := repositories.NewNewsRepository(db)
	topicRepository := repositories.NewTopicRepository(db)
	newsTopicRepository := repositories.NewNewsTopicRepository(db)

	newsService := services.NewNewsService(newsRepository, newsTopicRepository)
	topicService := services.NewTopicService(topicRepository)

	handlers.NewNewsHandler(newsService).Register(mux)
	handlers.NewTopicHandler(topicService).Register(mux)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
