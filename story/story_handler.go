package story

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"storyapi/utils"
	"strconv"

	"github.com/go-chi/chi"
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

//AddStory : to post stories
func (sh *Handler) AddStory(w http.ResponseWriter, r *http.Request) {
	log.Println("App : POST /add API hit!")
	var word AddStoryRequest
	body := json.NewDecoder(r.Body)
	err := body.Decode(&word)
	if err != nil {
		log.Println("JSON Decode Error in AddStory : ", err.Error())
		utils.Fail(w, utils.BadRequest, err.Error())
		return
	}
	err = ValidatePost(word)
	if err != nil {
		utils.Fail(w, utils.BadRequest, err.Error())
	}
	if word.Word != "" {
		response, status, err := sh.storyService.AddStory(&word)
		if err != nil {
			utils.Fail(w, status, err.Error())
		}
		utils.Send(w, status, response)
	} else {
		utils.Fail(w, utils.BadRequest, errors.New(utils.BadRequestError).Error())
	}
}

//GetStoryList : GET /stories endpoint
func (sh *Handler) GetStoryList(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /stories API hit")
	req := GetStoryRequest{
		Limit:  utils.DefaultLimit,
		Offset: utils.DefaultOffset,
		Sort:   utils.SortBy,
		Order:  utils.DESC,
	}
	request, err := validateGetStories(&req, r)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, utils.BadRequest, err.Error())
		return
	}
	stories, err := sh.storyService.GetStoryList(request)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, utils.InternalServerError, err.Error())
		return
	}
	log.Println("App : Stories listed successfully!")
	utils.Send(w, utils.Success, stories)
}

//GetStory : GET /stories/{id} endpoint
func (sh *Handler) GetStory(w http.ResponseWriter, r *http.Request) {
	log.Println("App : GET /stories/{id} API hit!")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, utils.InternalServerError, errors.New("invalid story id").Error())
		return
	}
	story, err := sh.storyService.GetStory(id)
	if err != nil {
		log.Println("Error : ", err.Error())
		utils.Fail(w, utils.InternalServerError, err.Error())
		return
	}
	log.Println("App : Story details fetched successfully")
	utils.Send(w, utils.Success, story)
}
