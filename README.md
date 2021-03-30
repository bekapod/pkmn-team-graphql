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

## Code Generation

To generate resolvers based on the schema:

```sh
go generate
```

To generate a dataloader for a type:

```sh
cd dataloader
go run github.com/vektah/dataloaden YourTypeLoader string "*bekapod/pkmn-team-graphql/data/model.YourType"
```

## Migrations

To generate a new database migration:

```sh
brew install golang-migrate
migrate create -ext sql -dir data/migrations -seq create_xxx_table
```

## Seed Data

To generate seed data:

```sh
LOG_LEVEL=debug go run scripts/seed/types/types.go
```

You can optionally specify a different output file:

```sh
LOG_LEVEL=debug go run scripts/seed/types/types.go --output-file seeds/something/some-file.sql
```

By default, data will be fetched directly from [https://pokeapi.co](https://pokeapi.co), however it is recommended to run PokeAPI locally to save hitting request limits. You can specify a custom PokeAPI host:

```sh
LOG_LEVEL=debug go run scripts/seed/types/types.go --pokeapi-host http://localhost
```

## Tests

Run tests:

```sh
go test ./...
```

Run tests with coverage:

```sh
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```
