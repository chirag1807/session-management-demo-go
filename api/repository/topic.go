package repository

import (
	"sessionmanagement/api/model/request"
	"sessionmanagement/api/model/response"
	"context"

	"github.com/jackc/pgx/v5"
)

type TopicRepostiry interface {
	AddTopic(request.Topic) error
	GetAllTopics() ([]response.Topic, error)
	UpdateTopic(request.Topic) error
	DeleteTopic(id int64) error
}

type topicRepostiry struct {
	pgx *pgx.Conn
}

func NewTopicRepo(pgx *pgx.Conn) TopicRepostiry {
	return topicRepostiry{
		pgx: pgx,
	}
}

func (t topicRepostiry) AddTopic(topic request.Topic) error {
	_, err := t.pgx.Exec(context.Background(), `INSERT INTO topics (name) VALUES ($1)`, topic.Name)
	if err != nil {
		return err
	}
	return nil
}

func (t topicRepostiry) GetAllTopics() ([]response.Topic, error) {
	topics, err := t.pgx.Query(context.Background(), `SELECT * FROM topics`)
	topicsSlice := make([]response.Topic, 0)

	if err != nil {
		return topicsSlice, err
	}
	defer topics.Close()

	var topic response.Topic
	for topics.Next() {
		if err := topics.Scan(&topic.ID, &topic.Name); err != nil {
			return topicsSlice, err
		}
		topicsSlice = append(topicsSlice, topic)
	}

	return topicsSlice, nil
}

func (t topicRepostiry) UpdateTopic(topic request.Topic) error {
	_, err := t.pgx.Exec(context.Background(), `UPDATE topics SET name = $1 WHERE id = $2`, topic.Name, topic.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t topicRepostiry) DeleteTopic(id int64) error {
	_, err := t.pgx.Exec(context.Background(), `DELETE FROM topics WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
