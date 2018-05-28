# Sample Go Restful API
This is an example showing how to create restful API with go and mysql

### Dependency
- `github.com/go-sql-driver/mysql` - mysql driver
- `github.com/gorilla/mux` - http router
- `github.com/Masterminds/squirrel` - mysql query builder
- `github.com/go-ini/ini` - ini config helper

### Installation
- Copy `config.sample.ini` to `config.ini` file and change configuration to match your environment
- This project use glide for handling dependency, read more on: `https://github.com/Masterminds/glide`.
```bash
glide install
go run main.go
```

# API specification
### Entities

##### News
```
type News struct {
	Id        *int   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

##### Topic
```
type Topic struct {
	Id        *int   `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### API endpoints

##### News

- `GET /news` - listing all news
- `GET /news/{id}` - get news by id
- `GET /news/{id}/topics` - get all topics of a news
- `GET /news/{id}/topics/{topicId}` - get 1 specific topic of a news by id
- `POST /news` - add a new news
- `PUT /news/{id}` - update a news
- `DELETE /news/{id}` - delete a news
- `POST /news/{id}/topics/{topicId}` - add a new relationship between news and topic (no need body)
- `DELETE /news/{id}/topics/{topicId}` - remove a relationship between news and topic

##### Topic

- `GET /topics` - listing all topics
- `GET /topics/{id}` - get topic by id
- `GET /topics/{id}/news` - get all news of a topic
- `GET /topics/{id}/news/{newsId}` - get 1 specific news of a topic by id
- `POST /topics` - add a new topics
- `PUT /topics/{id}` - update a topic
- `DELETE /topics/{id}` - delete a topic
- `POST /topics/{id}/news/{newsId}` - add a new relationship between news and topic (no need body)
- `DELETE /topics/{id}/news/{newsId}` - remove a relationship between news and topic