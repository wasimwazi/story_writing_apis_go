-- +goose Up
CREATE TABLE story (
    id SERIAL NOT NULL ,
    title character varying(100),
    created_at timestamp NOT NULL,
    updated_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE words (
    id SERIAL NOT NULL, 
    story_id INT, word varchar(50), 
    sentence_number INT, 
    para_number INT, 
    PRIMARY KEY(id), 
    FOREIGN KEY(story_id) REFERENCES story(id)
);

-- +goose Down
DROP TABLE words;
DROP TABLE story;
