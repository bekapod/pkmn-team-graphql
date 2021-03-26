# Pkmn Team - GraphQL API

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.16 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
git clone https://github.com/bekapod/pkmn-team-graphql.git
cd pkmn-team-graphql
go build -o bin/pkmn-team-graphql -v
heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

```sh
git push heroku main
heroku open
```

## Migrations

To generate a new database migration:

```sh
brew install golang-migrate
migrate create -ext sql -dir data/migrations -seq create_xxx_table
```
