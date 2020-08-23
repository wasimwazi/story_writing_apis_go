package story

import (
	"database/sql"
	"fmt"
	"log"
	"storyapi/utils"
	"strings"
)

//ServiceInterface : Story Service Interface
type ServiceInterface interface {
	AddStory(request *AddStoryRequest) (*Story, int, error)
}

//Service : Story Service Struct
type Service struct {
	storyRepo Repo
}

// NewService : Returns Story Service
func NewService(db *sql.DB) ServiceInterface {
	return &Service{
		storyRepo: NewRepo(db),
	}
}

//AddStory : Add story post request
func (service *Service) AddStory(request *AddStoryRequest) (*Story, int, error) {
	var story Story
	wordcount, err := service.storyRepo.GetLatestStory()
	if err != nil {
		return nil, 500, err
	}
	story.ID = wordcount.StoryID
	if wordcount.StoryID == 0 || wordcount.WordCount == utils.MaxStoryWordCount {
		story, err := service.storyRepo.CreateNewStory(request.Word)
		if err != nil {
			return nil, 500, err
		}
		log.Println("New story:", story)
		return story, 201, nil
	}
	if wordcount.WordCount == 0 {
		if len(strings.Split(wordcount.StoryTitle, " ")) == 1 {
			word := fmt.Sprintf("%s %s", wordcount.StoryTitle, request.Word)
			err = service.storyRepo.UpdateStoryTitle(wordcount.StoryID, word)
			if err != nil {
				return nil, 500, err
			}
			story.Title = word
			log.Println("Update story title:", story)
			return &story, 200, nil
		}
	}
	sentenceNumber := (wordcount.WordCount / 15) + 1
	paraNumber := (wordcount.WordCount / 150) + 1
	err = service.storyRepo.UpdateStoryWord(wordcount.StoryID, request.Word, sentenceNumber, paraNumber)
	if err != nil {
		return nil, 500, err
	}
	currentSentence, err := service.storyRepo.GetCurrentSentence(story.ID, sentenceNumber)
	if err != nil {
		return nil, 500, err
	}
	story.CurrentSentence = strings.Join(currentSentence, " ")
	story.Title = wordcount.StoryTitle
	log.Println("Update the sentence:", story)
	return &story, 200, nil
}
