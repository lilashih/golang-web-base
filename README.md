# Golang Web Base (Gin)

## ▍主要功能
本專案是一個以 `Gin` 為基礎的 `Golang` 網站，包含前後端，以下為常用指令，詳細開發文件請參閱 [document](/docs/document/)。


## ▍初次安裝
1. 建立環境設定檔
    ```
    cp .env.dev .env
    ```

2. 安裝相依套件並執行
    ```
    go mod tidy
    go run main.go migrate
    go run main.go
    ```

## ▍啟動
- 執行以下指令
    ```
    go run main.go
    ```

- 啟用熱加載
    1. 安裝 [Air](https://github.com/cosmtrek/air)  
        ```
        go install github.com/air-verse/air@latest
        ```
    2. 執行熱加載
        ```
        air
        ```

## ▍Migrate
```
go run main.go migrate
```

## ▍Swagger 文件
1. 安裝 [gin-swagger](https://github.com/swaggo/gin-swagger)
   
2. 重新產生 `API` 文件
    ```
    swag init
    ```
3. 在非 `release` 模式下才能開啟開啟 `/swagger/index.html`


## ▍測試
測試檔案攤平放在 ``test`` 目錄下，不要建立子目錄，然後執行
```
go test -v ./test/...
```

## ▍打包
### windows
- 一般打包
    ```
    go build -o release/app.exe
    ```

- 打包背景執行版本
    ```
    go build -ldflags="-H windowsgui" -o release/app.exe
    ```

    執行
    ```
    start /b app.exe
    ```

    結束
    ```
    taskkill /IM app.exe /F
    ```

- 跨平台打包（於 Docker 內執行）
    ```
    env GOOS=windows GOARCH=amd64 go build -o release/app
    ```

### macOS
```
env GOOS=darwin GOARCH=amd64 go build -o release/app_mac
```