package story

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"storyapi/utils"
)

//HandlerInterface : Story Handler Interface
type HandlerInterface interface {
	AddStory(w http.ResponseWriter, r *http.Request)
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
		log.Println("Error : ", err.Error())
		utils.Fail(w, 400, err.Error())
		return
	}
	err = ValidatePost(word)
	if err != nil {
		utils.Fail(w, 400, err.Error())
	}
	response, status, err := sh.storyService.AddStory(&word)
	if err != nil {
		utils.Fail(w, status, err.Error())
	}
	utils.Send(w, status, response)
}
