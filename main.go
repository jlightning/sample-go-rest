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
	"github.com/go-ini/ini"
)

type config struct {
	server struct {
		port string
	}
	database struct {
		host     string
		port     string
		database string
		username string
		password string
	}
}

func main() {
	cfg := loadConfig()
	dbConfig := cfg.database

	db, err := sql.Open("mysql", dbConfig.username+":"+dbConfig.password+"@tcp("+dbConfig.host+":"+dbConfig.port+")/"+dbConfig.database+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter().StrictSlash(true)

	newsRepository := repositories.NewNewsRepository(db)
	topicRepository := repositories.NewTopicRepository(db)
	newsTopicRepository := repositories.NewNewsTopicRepository(db)

	newsService := services.NewNewsService(newsRepository, newsTopicRepository)
	topicService := services.NewTopicService(topicRepository)

	handlers.NewNewsHandler(newsService, topicService).RegisterToRouter(router)
	handlers.NewTopicHandler(topicService, newsService).RegisterToRouter(router)

	log.Println("Listening...")
	http.ListenAndServe(":"+cfg.server.port, router)
}

func loadConfig() config {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}

	var configData config
	configData.server.port = cfg.Section("server").Key("port").String()

	configData.database.host = cfg.Section("database").Key("host").String()
	configData.database.port = cfg.Section("database").Key("port").String()
	configData.database.database = cfg.Section("database").Key("database").String()
	configData.database.username = cfg.Section("database").Key("username").String()
	configData.database.password = cfg.Section("database").Key("password").String()

	return configData
}
