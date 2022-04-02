# curso-go
Repositório do curso introdutório da linguagem Go

## Inicializando o módulo

    cd curso-go/
    go mod init github.com/eminetto/curso-go

## Executando

    go run main.go

## Compilando

```
go build -o bin/curso main.go
GOOS=darwin GOARCH=amd64 go build -o bin/curso main.go
GOOS=darwin GOARCH=arm64 go build -o bin/curso main.go
GOOS=windows GOARCH=amd64 go build -o bin/curso main.go
GOOS=linux GOARCH=amd64 go build -o bin/curso main.go
```