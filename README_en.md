# Golang Web Base (Gin)

## ▍Features
This project is a `Golang` web application based on `Gin`, including both frontend and backend. Below are common commands. For detailed development docs, see the [documentation](/docs/document/).


## ▍Installation
1. Create environment config file
    ```
    cp .env.dev .env
    ```

2. Install dependencies and run
    ```
    go mod tidy
    go run main.go migrate
    go run main.go
    ```

## ▍Start the server
- Run:
    ```
    go run main.go
    ```

- Enable hot reload (Air)
    1. Install [Air](https://github.com/cosmtrek/air)
        ```
        go install github.com/cosmtrek/air@latest
        ```
    2. Start hot reload
        ```
        air
        ```

## ▍Migrate
```
go run main.go migrate
```

## ▍Swagger 
1. Install [gin-swagger](https://github.com/swaggo/gin-swagger)

2. Regenerate `API` docs
    ```
    swag init
    ```


## ▍Testing
Place test files directly under the `test` directory (do not create subdirectories), then run:
```
go test -v ./test/...
```

## ▍Build
### Windows
- Regular build
    ```
    go build -o release/app.exe
    ```

- Build background (no console) version
    ```
    go build -ldflags="-H windowsgui" -o release/app.exe
    ```

    Run
    ```
    start /b app.exe
    ```

    Stop
    ```
    taskkill /IM app.exe /F
    ```

- Cross-platform build (run inside Docker)
    ```
    env GOOS=windows GOARCH=amd64 go build -o release/app
    ```

### macOS
```
env GOOS=darwin GOARCH=amd64 go build -o release/app_mac
```

