# STORY API
    
    APIs for collaborative story writing

## To setup the DB
    $ cd migration
    $ goose postgres <DB_URL> up
    $ cd ..

## To run the project

    Rename env.example to production.env
    Add the necessary environment variables
    $ source config/production.env
    $ go run main.go