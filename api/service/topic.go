package service

import (
	"sessionmanagement/api/model/request"
	"sessionmanagement/api/model/response"
	"sessionmanagement/api/repository"
)

type TopicService interface {
	AddTopic(request.Topic) error
	GetAllTopics() ([]response.Topic, error)
	UpdateTopic(request.Topic) error
	DeleteTopic(id int64) error
}

type topicService struct {
	topicRepostiry repository.TopicRepostiry
}

func NewTopicService(t repository.TopicRepostiry) TopicService {
	return topicService{
		topicRepostiry: t,
	}
}

func (t topicService) AddTopic(topic request.Topic) error {
	return t.topicRepostiry.AddTopic(topic)
}

func (t topicService) GetAllTopics() ([]response.Topic, error) {
	return t.topicRepostiry.GetAllTopics()
}

func (t topicService) UpdateTopic(topic request.Topic) error {
	return t.topicRepostiry.UpdateTopic(topic)
}

func (t topicService) DeleteTopic(id int64) error {
	return t.topicRepostiry.DeleteTopic(id)
}
