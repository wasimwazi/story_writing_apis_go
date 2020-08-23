package story

import (
	"database/sql"
)

// PostgresRepo : User Repo Struct for Postgres
type PostgresRepo struct {
	DB *sql.DB
}

//GetLatestStory : Postgres function to add a story
func (pg *PostgresRepo) GetLatestStory() (*WordCount, error) {
	var wordcount WordCount
	query := `
		SELECT 
			id, title 
		FROM 
			story 
		ORDER BY 
			id DESC 
		LIMIT 1
	`
	err := pg.DB.QueryRow(query).Scan(&wordcount.StoryID, &wordcount.StoryTitle)
	if err == sql.ErrNoRows {
		return &wordcount, nil
	}
	if err != nil {
		return nil, err
	}
	query = `
		SELECT 
			COUNT(*)
		FROM
			words
		WHERE
			story_id = $1
	`
	err = pg.DB.QueryRow(query, wordcount.StoryID).Scan(&wordcount.WordCount)
	if err != nil {
		return nil, err
	}
	return &wordcount, nil
}

//CreateNewStory : Create a new story in story table
func (pg *PostgresRepo) CreateNewStory(word string) (*Story, error) {
	query := `
		INSERT INTO story (title, created_at, updated_at) 
		VALUES ($1, NOW(), NOW()) RETURNING id
	`
	var story Story
	err := pg.DB.QueryRow(query, word).Scan(&story.ID)
	if err != nil {
		return nil, err
	}
	story.Title = word
	return &story, nil
}

//UpdateStoryTitle : To update the title of story
func (pg *PostgresRepo) UpdateStoryTitle(id int, word string) error {
	query := `
		UPDATE 
			story
		SET
			title = $1, 
			updated_at = NOW()
		WHERE
			id = $2 
	`
	_, err := pg.DB.Exec(query, word, id)
	return err
}

//UpdateStoryWord : Update the words in story
func (pg *PostgresRepo) UpdateStoryWord(id int, word string, sentenceNumber int, paraNumber int) error {
	tx, err := pg.DB.Begin()
	if err != nil {
		return err
	}
	query := `
		INSERT INTO words (word, sentence_number, para_number, story_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err = tx.Exec(query, word, sentenceNumber, paraNumber, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = `
		UPDATE
			story
		SET
			updated_at = NOW()
		WHERE
			id = $1
	`
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

//GetCurrentSentence : To get the current sentence
func (pg *PostgresRepo) GetCurrentSentence(id int, sentenceNumber int) ([]string, error) {
	var sentence []string
	query := `
		SELECT 
			word
		FROM 
			words
		WHERE 
			story_id = $1
		AND
			sentence_number = $2
		ORDER BY
			id ASC 
	`
	rows, err := pg.DB.Query(query, id, sentenceNumber)
	if err != nil {
		return nil, err
	}
	var word string
	for rows.Next() {
		err := rows.Scan(&word)
		if err != nil {
			return nil, err
		}
		sentence = append(sentence, word)
	}
	return sentence, nil
}
