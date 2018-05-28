package entities

type NewsTopic struct {
	Id      *int `json:"id"`
	NewsId  int  `json:"new_id"`
	TopicId int  `json:"topic_id"`
}
