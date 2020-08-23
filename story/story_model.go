package story

import "time"

//AddStoryRequest : Story Request Struct
type AddStoryRequest struct {
	Word string `json:"word"`
}

//Story :: Struct to represent a story
type Story struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	CurrentSentence string `json:"current_sentence"`
}

//WordCount : To get the story word count
type WordCount struct {
	StoryID    int
	StoryTitle string
	WordCount  int
}

//GetStoryRequest : Story List Request Struct
type GetStoryRequest struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	Sort string `json:"sort"`
	Order string `json:"order"`
}

//SingleStory : Single story struct
type SingleStory struct {
	ID int `json:"id"`
	Title string `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//StoryList : Story listing Struct
type StoryList struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	Count int `json:"count"`
	Results []SingleStory `json:"results"`
}
