# Greenlight API with Go

Greenlight is a Rest API project of the `Alex Edwards` book. I'm using this book for learning how to build Rest APIs. 

- https://lets-go-further.alexedwards.net/

## Enable modules

```
go mod init github.com/williammunozr/greenlight
```

## Dependencies

```
go get github.com/julienschmidt/httprouter@v1.3.0
```

## Start the PostgreSQL Database

```
sudo service postgresql start
```

## Run the API

```
go run ./cmd/api
```