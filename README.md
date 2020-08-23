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