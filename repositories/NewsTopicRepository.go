package repositories

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type INewsTopicRepository interface {
	InsertRelationship(newsId int, topicId int) error
	DeleteRelationship(newsId int, topicId int) error
}

type newsTopicRepository struct {
	db *sql.DB
}

func NewNewsTopicRepository(db *sql.DB) *newsTopicRepository {
	return &newsTopicRepository{db: db}
}

func (repository *newsTopicRepository) InsertRelationship(newsId int, topicId int) error {
	sql, args, err := squirrel.Insert("news_topic").Columns("news_id", "topic_id").Values(newsId, topicId).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(sql, args...)
	return err
}

func (repository *newsTopicRepository) DeleteRelationship(newsId int, topicId int) error {
	sql, args, err := squirrel.Delete("news_topic").Where(squirrel.Eq{"news_id": newsId, "topic_id": topicId}).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(sql, args...)
	return err
}
