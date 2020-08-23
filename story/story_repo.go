package story

import "database/sql"

//Repo : Story Repository Interface
type Repo interface {
	GetLatestStory() (*WordCount, error)
	CreateNewStory(string) (*Story, error)
	UpdateStoryTitle(int, string) error
	UpdateStoryWord(int, string, int, int) error
	GetCurrentSentence(int, int) ([]string, error)
	GetStoryList(*GetStoryRequest) ([]SingleStory, error)
	GetStory(int) (*SingleStory, error)
	GetWordsInStory(int) ([]Words, error)
}

// NewRepo : Returns Story Repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
