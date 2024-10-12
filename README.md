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

```shell
brew install golang-migrate
```

### Validate the migrate installation

```
migrate -version
v4.18.1
```

## Start the PostgreSQL Database

In the new environment I have Docker Desktop with a PostgreSQL Docker image, just open Docker Desktop and start the container.
To admin the database I'm using the application DBeaver.

ToDo: 
- Documentation at [Rust Newsletter App](https://github.com/therustdeveloper/newsletter)

## Install psql

```shell
brew install libpq
```

## Link psql

```shell
brew link --force libpq
```

## Connect to PostgreSQL Database

```commandline
psql postgres -h localhost -U postgres
Password for user postgres: password
```

## Create Database

```commandline
postgres=# create database greenlight;
CREATE DATABASE
```

## Create User

```commandline
postgres=# create user greenlight with encrypted password 'metallica';
CREATE ROLE
```

## Assign Permissions to User on the Database

```commandline
postgres=# grant all privileges on database greenlight to greenlight;
GRANT
```

## Connect to greenlight Database

```commandline
psql greenlight -h localhost -U greenlight
Password for user postgres: metallica
```

## Grant all on public to greenlight

```commandline
newsletter=# grant all on schema public to greenlight;
GRANT
newsletter=#
```

## Connect to PostgreSQL Database

```commandline
psql postgres -h localhost -U postgres
Password for user postgres: password
```

## Create extension

```commandline
postgres=# CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION
```

## Allow create databases

```commandline
postgres=# alter user greenlight createdb;
ALTER ROLE
```

## Grant permissions to greenlight user

```commandline
postgres=# alter role greenlight superuser;
ALTER ROLE
```

## Create GREENLIGHT_DB_DSN environment variable

```commandline
export GREENLIGHT_DB_DSN=postgres://greenlight:metallica@localhost/greenlight?sslmode=disable
```

## Create citext extension in the greenlight database

```shell
psql $GREENLIGHT_DB_DSN 
psql (14.13 (Homebrew), server 15.3 (Debian 15.3-1.pgdg110+1))
WARNING: psql major version 14, server major version 15.
         Some psql features might not work.
Type "help" for help.

greenlight=# CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION
```

## Run the migration

```shell
migrate -path migrations -database $GREENLIGHT_DB_DSN up
1/u create_movies_table (8.130041ms)
2/u add_movies_check_constraints (12.553875ms)
3/u add_movies_indexes (16.790958ms)
4/u create_users_table (25.081666ms)
5/u create_tokens_table (30.3545ms)
6/u add_permissions (38.714ms)
```

## Run the API

```
go run ./cmd/api -db-dsn=$GREENLIGHT_DB_DSN
{"level":"INFO","time":"2023-08-19T17:04:05Z","message":"database connection pool established"}
{"level":"INFO","time":"2023-08-19T17:04:05Z","message":"starting server","properties":{"addr":":4000","env":"development"}}
```

## Stop the API

```commandline
^C{"level":"INFO","time":"2024-10-12T16:45:56Z","message":"shutting down server","properties":{"signal":"interrupt"}}
{"level":"INFO","time":"2024-10-12T16:45:56Z","message":"completing background task","properties":{"addr":":4000"}}
{"level":"INFO","time":"2024-10-12T16:45:56Z","message":"stopped server","properties":{"addr":":4000"}}
```