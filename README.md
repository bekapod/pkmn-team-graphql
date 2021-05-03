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

To generate resolvers based on the GraphQL schema (`graph/schema.graphqls`):

```sh
go run github.com/99designs/gqlgen generate
```

To generate a dataloader for a type:

```sh
cd dataloader
go run github.com/vektah/dataloaden TypeLoader string "*bekapod/pkmn-team-graphql/data/model.Type"
```

## Migrations

To generate a new database migration, you should first update `data/schema.prisma` and then run:

```sh
go run github.com/prisma/prisma-client-go migrate dev --schema data/schema.prisma --preview-feature
```

## Seed Data

To generate seed data:

```sh
LOG_LEVEL=debug go run scripts/seed/types/types.go
```

By default, data will be fetched directly from [https://pokeapi.co](https://pokeapi.co), however it is recommended to run PokeAPI locally to save hitting request limits. You can specify a custom PokeAPI host:

```sh
LOG_LEVEL=debug go run scripts/seed/types/types.go --pokeapi-host http://localhost
```

## Tests

### Unit Tests

Run tests:

```sh
go test ./...
```

Run tests with coverage:

```sh
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

### API Tests

```sh
yarn global add newman
newman run tests/Pokemon\ Data.postman_collection.json -e tests/Local.postman_environment.json
```

```sh
brew install k6
k6 run tests/pokemon-data-k6.js
```
