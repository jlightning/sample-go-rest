package repositories

import (
	"database/sql"
	"sample-go-rest/entities"
	"github.com/Masterminds/squirrel"
	"errors"
)

type INewsRepository interface {
	GetList() ([]entities.News, error)
	GetItemById(id int) (*entities.News, error)
	InsertItem(news entities.News) error
	UpdateItem(id int, news entities.News) error
	DeleteItem(id int) error
}

type newsRepostory struct {
	db *sql.DB
}

func NewNewsRepository(db *sql.DB) INewsRepository {
	return &newsRepostory{db: db}
}

func (repository *newsRepostory) GetList() ([]entities.News, error) {
	sql, args, err := squirrel.Select("*").From("news").ToSql()
	if err != nil {
		return nil, err
	}
	result, err := repository.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	list := []entities.News{}

	for result.Next() {
		var news entities.News
		result.Scan(&news.Id, &news.Title, &news.Content)

		list = append(list, news)
	}

	return list, nil
}

func (repository *newsRepostory) GetItemById(id int) (*entities.News, error) {
	sql, args, err := squirrel.Select("*").From("news").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}
	result, err := repository.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var news entities.News
		result.Scan(&news.Id, &news.Title, &news.Content)

		return &news, nil
	}

	return nil, errors.New("item not found")
}

func (repository *newsRepostory) InsertItem(news entities.News) error {
	sql, args, err := squirrel.Insert("news").Columns("title", "content").Values(news.Title, news.Content).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(sql, args...)
	return err
}

func (repository *newsRepostory) UpdateItem(id int, news entities.News) error {
	sql, args, err := squirrel.Update("news").SetMap(map[string]interface{}{"title": news.Title, "content": news.Content}).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(sql, args...)
	return err
}

func (repository *newsRepostory) DeleteItem(id int) error {
	sql, args, err := squirrel.Delete("news").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(sql, args...)
	return err
}
