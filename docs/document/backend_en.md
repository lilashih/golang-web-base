# Golang Web Base (Gin) — Backend Development

## ▍Packages

| Package | Description |
| :----------- | :----------- |
| [gin-gonic/gin](https://github.com/gin-gonic/gin) | High-performance HTTP framework; handles routing and middleware. |
| [gin-contrib/cors](https://github.com/gin-contrib/cors) | CORS middleware for Gin, used for cross-origin settings. |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | ORM framework for database models and operations. |
| [gorm.io/driver/sqlite](https://github.com/go-gorm/sqlite) | SQLite driver. |
| [swaggo/swag](https://github.com/swaggo/swag) | Generates Swagger docs from annotations. |
| [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger) | Integrates Swagger UI into Gin. |
| [go-playground/validator](https://github.com/go-playground/validator) | Struct and field validation. |
| [go-playground/universal-translator](https://github.com/go-playground/universal-translator) | i18n translations for validation errors. |
| [go-playground/locales](https://github.com/go-playground/locales) | Provides multilingual locale data. |
| [caarlos0/env](https://github.com/caarlos0/env) | Parses environment variables into structs. |
| [golobby/dotenv](https://github.com/golobby/dotenv) | Loads `.env` configuration files. |
| [jinzhu/copier](https://github.com/jinzhu/copier) | Copies struct fields (DTO ↔ Model conversion). |

## ▍Directory Structure
```
.
├── .air.toml
├── .env.example
├── go.mod
├── go.sum
├── main.go
├── docs/                    # Swagger and other documentation
├── release/                 # Build output for deployment, includes frontend files
|   └── public/              # Frontend app
├── src/                     # Backend source code
|   ├── core/                # Core modules
|   |   ├── config/          # Environment configuration
|   |   ├── db/              # Database connection
|   |   ├── helper/          # Utilities
|   |   ├── http/
|   |   |   ├── request/     # Request types
|   |   |   └── response/    # Response types
|   |   ├── logger/          # System logging
|   |   ├── migrate/         # Migration core
|   |   └── model/           # Base model
|   ├── def/                 # Definitions
|   ├── http/
|   |   ├── controller/      # Controllers
|   |   └── resource/        # API response wrappers
|   ├── migrate/             # Migration scripts
|   ├── model/               # GORM ORM models
|   ├── repository/          # Data access layer
|   └── route/               # Routes
├── storage/                 # Local storage
└── test/                    # Unit tests; see test.md
```

## ▍Environment Variables
All settings are located in [src/core/config](/src/core/config), automatically initialized by [src/core/config/main.go](/src/core/config/main.go), and loaded from `.env`.

```go
import (
  "fmt"
  "gbase/src/core/config"
)

fmt.Println(config.App.Name)
```

### Common .env Variables

| Name | Purpose | Default (production) | Notes |
|---|---|---|---|
| `APP_PORT` | HTTP server port | `8000` |  |
| `APP_MODE` | Run mode (`release` / `debug` / `test`) | `release` | Use `release` in production. |
| `APP_PUBLIC_PATH` | Frontend static file path | `./public` | Usually `./release/public` in development. |
| `DB_MAIN_PATH` | Primary DB path | `./database/db.sqlite` | SQLite path; usually `./storage/database/db.sqlite` in development. |
| `DB_SECOND_PATH` | Secondary DB path | `./database/db_second.sqlite` |  |

### Security Notes
- Do not commit `.env` to version control (add to `.gitignore`).
- `APP_MODE=debug` is for development only.

## ▍Logging
| Logger | Purpose | Filename | Behavior by environment |
|---|---|---|---|
| `logger.Log` | General system logs | `log-YYYY-MM-DD.log` | In `release` mode writes to file only; other modes also print to console. |
| `logger.Migrate` | Migration-specific logs | `migrate-YYYY-MM-DD.log` | In `test` mode logs are discarded (`io.Discard`). |

- File path: logs are stored under `storage/log` at the project root.  
- Features:  
  - New file per day (with UTF-8 BOM).
  - Customizable format, outputs, and rotation in [src/core/logger](/src/core/logger).

```go
import (
  "gbase/src/core/logger"
)

// Write general log
logger.Log.Println("Server started")

// Write migration log
logger.Migrate.Println("Applying migrations...")
```


## ▍Migration
- Migration files: [src/migrate/migration](/src/migrate/migration/).
- Entrypoint: [src/migrate/main.go](/src/migrate/main.go).
- Run:  
    ```bash
    go run main.go migrate
    ```
- Output logs: `/storage/log/migrate-YYYY-MM-DD.log`.
- No rollback mechanism.


## ▍Routing
### Structure
- `RESTful API` style.
- Routes defined in [src/route/main.go](/src/route/main.go).
- Controller logic in [src/http/controller](/src/http/controller/).

### Static Assets & SPA
| Path | Maps to | Notes |
|-----------|-----------|-----------|
| `/static` | `<PublicPath>/static` | |
| `/src` | `<PublicPath>/app/src` | |
| `/app` | `<PublicPath>/app/index.html` | SPA entry point. Any route that starts with `/app` will serve `index.html` for client-side routing. |
| `/` |  | Redirects to `/app`。 |

- `PublicPath` is set by `config.App.PublicPath`.

### API Registration Example
```go
api := router.Group("/api")

controller.NewAppController().RegisterRoutes(api)
controller.NewUserController().RegisterRoutes(api)
```

### Swagger Docs
- Enabled only when not in `release` mode.
- Default path: `/swagger/index.html`.

## ▍API Documentation Generation (Swagger)
Generate `API` docs compliant with the `OpenAPI` specification — [OpenAPI Specification (OAS)](https://swagger.io/specification/).

1. Install [gin-swagger](https://github.com/swaggo/gin-swagger).  
2. Add API annotations under [src/http/controller](/src/http/controller/). See the [swag README](https://github.com/swaggo/swag/blob/master/README_zh-CN.md) or the [OpenAPI spec](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md) for details.
3. Generate docs:  
  ```shell
  swag init
  ```

