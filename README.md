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
go install honnef.co/go/tools/cmd/staticcheck@latest
go get github.com/tomasen/realip@latest
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

```
sudo service postgresql start
```

## Run the API

```
go run ./cmd/api
```