# CHI backend 2 tiers APP

## Roadmap

- [x] Migration
- [x] Seeding
- [x] Router
- [ ] Auth (En cours)
- [ ] 0auth
- [ ] OpenAPI
- [ ] Validator

## Install

- Docker
- Goose
- Sqlc

### Commands

Run server

`go run main.go`

Create migration

`goose -dir ./migrations create <changeme> sql`

Apply migration

`goose -dir ./migrations postgres "postgres://postgres:password@localhost:5432/postgres" up`

Downgrade migration

`goose -dir ./migrations postgres "postgres://postgres:password@localhost:5432/postgres" down`

Seeding

`cd ./seeds`
`goose create <changeme> go`
`go build -o goose-seed *.go`
`./goose-seed postgres "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" up -no-versioning`
    
### TODO

- Move db string to env vars 
