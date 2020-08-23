package story

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
