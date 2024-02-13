package controller

import (
	"encoding/json"
	"net/http"
	"sessionmanagement/api/model/request"
	"sessionmanagement/api/model/response"
	"sessionmanagement/api/service"
	"sessionmanagement/constants"
	"sessionmanagement/error"
	"sessionmanagement/utils"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type TopicController interface {
	AddTopic(w http.ResponseWriter, r *http.Request)
	GetAllTopics(w http.ResponseWriter, r *http.Request)
	UpdateTopic(w http.ResponseWriter, r *http.Request)
	DeleteTopic(w http.ResponseWriter, r *http.Request)
}

type topicController struct {
	topicService service.TopicService
	sessionManager *scs.SessionManager
}

func NewTopicController(t service.TopicService, sessionManager *scs.SessionManager) TopicController {
	return topicController{
		topicService: t,
		sessionManager: sessionManager,
	}
}

func (t topicController) AddTopic(w http.ResponseWriter, r *http.Request) {
	var topic request.Topic
	err := json.NewDecoder(r.Body).Decode(&topic)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = t.topicService.AddTopic(topic)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.TOPIC_ADDED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
	return
}

func (t topicController) GetAllTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := t.topicService.GetAllTopics()

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.TopicResponse{
		Topics: topics,
	}
	utils.ResponseGenerator(w, 200, response)
	return
}

func (t topicController) UpdateTopic(w http.ResponseWriter, r *http.Request) {
	var topic request.Topic
	err := json.NewDecoder(r.Body).Decode(&topic)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = t.topicService.UpdateTopic(topic)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.TOPIC_UPDATED,
	}
	utils.ResponseGenerator(w, 200, response)
	return
}

func (t topicController) DeleteTopic(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	err := t.topicService.DeleteTopic(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.TOPIC_DELETED,
	}
	utils.ResponseGenerator(w, 200, response)
	return
}
