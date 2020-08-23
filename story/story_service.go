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
	GetStoryList(request *GetStoryRequest) (*StoryList, error)
	GetStory(id int) (*StoryData, error)
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
		return nil, utils.InternalServerError, err
	}
	story.ID = wordcount.StoryID
	if wordcount.StoryID == 0 || wordcount.WordCount == utils.MaxStoryWordCount {
		story, err := service.storyRepo.CreateNewStory(request.Word)
		if err != nil {
			return nil, utils.InternalServerError, err
		}
		return story, utils.Created, nil
	}
	if wordcount.WordCount == 0 {
		if len(strings.Split(wordcount.StoryTitle, " ")) == 1 {
			word := fmt.Sprintf("%s %s", wordcount.StoryTitle, request.Word)
			err = service.storyRepo.UpdateStoryTitle(wordcount.StoryID, word)
			if err != nil {
				return nil, utils.InternalServerError, err
			}
			story.Title = word
			return &story, utils.Success, nil
		}
	}
	sentenceNumber := (wordcount.WordCount / utils.SentenceLength) + 1
	paraNumber := (wordcount.WordCount / utils.WordsInParagraph) + 1
	err = service.storyRepo.UpdateStoryWord(wordcount.StoryID, request.Word, sentenceNumber, paraNumber)
	if err != nil {
		return nil, utils.InternalServerError, err
	}
	currentSentence, err := service.storyRepo.GetCurrentSentence(story.ID, sentenceNumber)
	if err != nil {
		return nil, utils.InternalServerError, err
	}
	story.CurrentSentence = strings.Join(currentSentence, " ")
	story.Title = wordcount.StoryTitle
	return &story, utils.Success, nil
}

//GetStoryList : Get the list of stories
func (service *Service) GetStoryList(request *GetStoryRequest) (*StoryList, error) {
	var storylist StoryList
	storylist.Limit = request.Limit
	storylist.Offset = request.Offset
	stories, err := service.storyRepo.GetStoryList(request)
	if err != nil {
		return nil, err
	}
	storylist.Count = len(stories)
	storylist.Results = stories
	return &storylist, nil
}

//GetStory : Get the details of a story
func (service *Service) GetStory(id int) (*StoryData, error) {
	var storydata StoryData
	story, err := service.storyRepo.GetStory(id)
	if err != nil {
		return nil, err
	}
	storydata.ID = story.ID
	storydata.Title = story.Title
	storydata.CreatedAt = story.CreatedAt
	storydata.UpdatedAt = story.UpdatedAt
	words, err := service.storyRepo.GetWordsInStory(id)
	if err != nil {
		return nil, err
	}
	var paraArray []Paragraph
	if len(words) > 0 {
		sentenceMap := make(map[SentencePara]Sentence)
		for _, v := range words {
			var sp SentencePara
			sp.ParaNumber = v.ParaNumber
			sp.SentenceNumber = v.SentenceNumber
			sentenceMap[sp] = append(sentenceMap[sp], v.Word)
		}
		log.Println(len(sentenceMap))
		sentenceStringMap := make(map[SentencePara]string)
		for k, v := range sentenceMap {
			sentenceStringMap[k] = strings.Join(v, " ")
		}
		paragraphSentence := make(map[int][]string)
		for k, v := range sentenceStringMap {
			paragraphSentence[k.ParaNumber] = append(paragraphSentence[k.ParaNumber], v)
		}
		paragraphMap := make(map[int]Paragraph)
		for k, v := range paragraphSentence {
			var p Paragraph
			p.Sentences = v
			paragraphMap[k] = p
		}
		for _, v := range paragraphMap {
			paraArray = append(paraArray, v)
		}
	}
	storydata.Paragraph = paraArray
	return &storydata, nil
}
