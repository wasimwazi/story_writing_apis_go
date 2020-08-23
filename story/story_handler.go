package story

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"storyapi/utils"
	"strconv"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

//HandlerInterface : Story Handler Interface
type HandlerInterface interface {
	AddStory(w http.ResponseWriter, r *http.Request)
	GetStoryList(w http.ResponseWriter, r *http.Request)
	GetStory(w http.ResponseWriter, r *http.Request)
}

//Handler : Story Handler Struct
type Handler struct {
	storyService ServiceInterface
}

//NewHTTPHandler : Returns Story HTTP Handler
func NewHTTPHandler(db *sql.DB) HandlerInterface {
	return &Handler{
		storyService: NewService(db),
	}
}

//AddStory : Handler for POST /add request
func (sh *Handler) AddStory(w http.ResponseWriter, r *http.Request) {
	log.WithFields(
		log.Fields{
			"Function": "AddStory()",
		}).Debug("App : POST /add API hit!")
	var word AddStoryRequest
	body := json.NewDecoder(r.Body)
	err := body.Decode(&word)
	if err != nil {
		log.WithFields(
			log.Fields{
				"Function": "AddStory()",
			}).Error("Error : JSON Decode Error : ", err.Error())
		utils.Fail(w, utils.BadRequest, err.Error())
		return
	}
	err = ValidatePost(word)
	if err != nil {
		log.WithFields(
			log.Fields{
				"Function": "AddStory()",
			}).Error(fmt.Errorf("Error : %w", err))
		utils.Fail(w, utils.BadRequest, err.Error())
		return
	}
	if word.Word != "" {
		response, status, err := sh.storyService.AddStory(&word)
		if err != nil {
			log.WithFields(
				log.Fields{
					"Function": "AddStory()",
				}).Error(fmt.Errorf("Error : Error while adding story - %w", err))
			utils.Fail(w, status, err.Error())
			return
		}
		log.WithFields(
			log.Fields{
				"Function": "GetStoryList()",
				"Word":     word.Word,
			}).Debug("App : Story updated successfully!")
		utils.Send(w, status, response)
		return
	}
	utils.Fail(w, utils.BadRequest, errors.New(utils.BadRequestError).Error())
	return
}

//GetStoryList : Handler for GET /stories request
func (sh *Handler) GetStoryList(w http.ResponseWriter, r *http.Request) {
	log.WithFields(
		log.Fields{
			"Function": "GetStoryList()",
		}).Debug("App : GET /stories API hit")
	req := GetStoryRequest{
		Limit:  utils.DefaultLimit,
		Offset: utils.DefaultOffset,
		Sort:   utils.SortBy,
		Order:  utils.DESC,
	}
	request, err := ValidateGetStories(&req, r)
	if err != nil {
		log.WithFields(
			log.Fields{
				"Function": "GetStoryList()",
			}).Error("Error : ", err.Error())
		utils.Fail(w, utils.BadRequest, err.Error())
		return
	}
	stories, err := sh.storyService.GetStoryList(request)
	if err != nil {
		log.WithFields(
			log.Fields{
				"Function": "GetStoryList()",
			}).Error("Error : Error while fetching fetching from DB ", err.Error())
		utils.Fail(w, utils.InternalServerError, err.Error())
		return
	}
	log.WithFields(
		log.Fields{
			"Function": "GetStoryList()",
			"Limit":    stories.Limit,
			"Offset":   stories.Offset,
			"SortBy":   req.Sort,
			"OrderBy":  req.Order,
		}).Debug("App : Stories listed successfully!")
	utils.Send(w, utils.Success, stories)
}

//GetStory : Handler for GET /stories/{id} request
func (sh *Handler) GetStory(w http.ResponseWriter, r *http.Request) {
	log.WithFields(
		log.Fields{
			"Function": "GetStory()",
		}).Debug("App : GET /stories/{id} API hit!")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.WithFields(
			log.Fields{
				"Function": "GetStory()",
			}).Error("Error : Error while converting string to integer ", err.Error())
		utils.Fail(w, utils.InternalServerError, errors.New("invalid story id").Error())
		return
	}
	story, err := sh.storyService.GetStory(id)
	if err != nil {
		log.WithFields(
			log.Fields{
				"Function": "GetStory()",
			}).Error("Error : Error while fetching from DB ", err.Error())
		utils.Fail(w, utils.InternalServerError, err.Error())
		return
	}
	log.WithFields(
		log.Fields{
			"Function": "GetStory()",
			"Story ID": id,
		}).Debug("App : Story details fetched successfully")
	utils.Send(w, utils.Success, story)
}
