# Greenlight API with Go

Greenlight is a Rest API project of the `Alex Edwards` book. I'm using this book for learning how to build Rest APIs. 

- https://lets-go-further.alexedwards.net/

## Enable modules

```
go mod init github.com/thegodeveloper/greenlight
```

## Dependencies

```
go get github.com/julienschmidt/httprouter@v1.3.0
go install honnef.co/go/tools/cmd/staticcheck@latest
go get github.com/tomasen/realip@latest
```

## Other Software

### Installing hey

```commandline
$ go install github.com/rakyll/hey@latest
```

### Installing the migrate tool

```
cd .\Downloads\
curl -o migrate.windows-amd64.zip https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.windows-amd64.zip
unzip migrate.windows-amd64.zip
mv .\migrate.exe ..\go\bin\
```

### Validate the migrate installation

```
migrate -version
4.15.2
```

## Start the PostgreSQL Database

In the new environment I have Docker Desktop with a PostgreSQL Docker image, just open Docker Desktop and start the container.
To admin the database I'm using the application DBeaver.

## Grant permissions to greenlight user

```shell
psql postgres

alter role greenlight superuser;
```

## Create citext extension in the greenlight database

```shell
psql $GREENLIGHT_DB_DSN 

CREATE EXTENSION IF NOT EXISTS citext;
```

## Run the migration

```shell
migrate -path migrations -database $GREENLIGHT_DB_DSN up
```

## Run the API

```
go run ./cmd/api -db-dsn=$GREENLIGHT_DB_DSN
{"level":"INFO","time":"2023-08-19T17:04:05Z","message":"database connection pool established"}
{"level":"INFO","time":"2023-08-19T17:04:05Z","message":"starting server","properties":{"addr":":4000","env":"development"}}
```