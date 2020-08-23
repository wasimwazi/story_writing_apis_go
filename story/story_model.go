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

//WordCount : Struct to represent story and word count
type WordCount struct {
	StoryID    int
	StoryTitle string
	WordCount  int
}

//GetStoryRequest : Story listing request Struct
type GetStoryRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
}

//SingleStory : Story struct to represent a single story
type SingleStory struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//StoryList : Stories listing Struct
type StoryList struct {
	Limit   int           `json:"limit"`
	Offset  int           `json:"offset"`
	Count   int           `json:"count"`
	Results []SingleStory `json:"results"`
}

//Sentence : To represent a single sentence
type Sentence []string

//SentencePara : Sentence and paragraph number struct
type SentencePara struct {
	SentenceNumber int
	ParaNumber     int
}

//Paragraph : To represent story paragraph struct
type Paragraph struct {
	Sentences []string `json:"sentences"`
}

//StoryData : To represent details of a story
type StoryData struct {
	ID        int         `json:"id"`
	Title     string      `json:"title"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Paragraph []Paragraph `json:"paragraph"`
}

//Words : Struct to represent a word
type Words struct {
	ID             int
	Word           string
	SentenceNumber int
	ParaNumber     int
}
