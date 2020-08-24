# STORY API
    
    APIs for collaborative story writing

## To run the project locally
### To setup the DB
    $ cd migration
    $ goose postgres <DB_URL> up
    $ cd ..

### To run the project

    Add the necessary environment variables to config/production.env
    $ source config/production.env
    $ go run main.go

## To run the project using Docker
    
    $ docker-compose up

    The server will be running on port 3000

## If you had more time, what would you do differently?
    1. Add options for story edit and delete options.
    2. Add functionality to maintain story edit history so that everyone can analyse the contributions in the story.