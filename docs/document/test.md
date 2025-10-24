# Golang Web Base (Gin) 單元測試

本專案的測試檔位於 `test/`，直接使用 `route.SetupRouter()` 建立一個 `Gin` 引擎並透過 `httptest` 發送請求（不會啟動實際網路埠），且測試時會使用獨立的測試資料庫，不會影響開發資料。  

測試入口（TestMain）及常用的 `helper` 在 `test/main_test.go`，包含 `List`、`Create`、`Update`、`Delete` 等函式，方便對 `API` 做整合式測試。

## ▍目錄架構
```
test/
├── .env            # 測試用環境設定，需加入版本控制
├── main_test.go    # 測試入口（TestMain）及 helper
├── storage/        # 用於放置測試期間產生的 SQLite 檔或其他臨時檔案，使用獨立路徑以免污染開發資料
└── xxx_test.go     # 對應資源的測試案例，例如 user_test.go 測試 users API (測試檔案請攤平放在 test 下，不要建立子目錄)
```


## ▍環境變數 

### 測試版本的 .env
| 名稱 | 測試值 | 備註 |
|---|---|---|
| `APP_MODE` | `test` | 必須設為 `test` 才會啟用測試行為
| `DB_MAIN_PATH` | `./storage/database/db.sqlite` | 指向專用測試 DB 檔，以免污染開發資料
| `DB_SECOND_PATH` | `./storage/database/second.sqlite` | 



## ▍執行
```shell
go test -v ./test/...
```