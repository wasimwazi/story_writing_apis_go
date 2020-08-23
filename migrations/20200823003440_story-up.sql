-- +goose Up
CREATE TABLE IF NOT EXISTS story (
    id SERIAL NOT NULL ,
    title character varying(100) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS words (
    id SERIAL NOT NULL, 
    story_id INT NOT NULL, 
    word varchar(50) NOT NULL, 
    sentence_number INT NOT NULL, 
    para_number INT NOT NULL, 
    PRIMARY KEY(id), 
    FOREIGN KEY(story_id) REFERENCES story(id)
);

-- +goose Down
DROP TABLE words;
DROP TABLE story;
