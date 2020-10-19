# Migrations

https://github.com/golang-migrate/migrate

## Create migration

```migrate create -ext sql -dir migrations <migration name>```

## UP migrations

```migrate -path migrations -database "postgres://postgres@localhost/recumbent_stone_dev" up```

## DOWN migrations

```migrate -path migrations -database "postgres://postgres@localhost/recumbent_stone_dev" down```