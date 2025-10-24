package route

import (
	"fmt"
	"gbase/src/core/config"
	"gbase/src/core/helper"
	"gbase/src/http/controller"
	"log"
	"path/filepath"

	docs "gbase/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title App
func StartService() {
	router := SetupRouter()

	routerErr := router.Run(fmt.Sprintf(":%d", config.App.Port))
	if routerErr != nil {
		log.Fatal(routerErr)
	}
}

func SetupRouter() *gin.Engine {
	gin.SetMode(config.App.Mode)
	router := gin.Default()
	// router.ForwardedByClientIP = false

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	router.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(204)
		}
	})

	// router.Use(middleware.AccessLog())

	// 前端SPA頁面：所有 / 開頭的都回傳 index.html (但避免 src 衝突)
	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(config.App.PublicPath, "app", "index.html"))
	})

	// 靜態資源
	router.Static("/static", filepath.Join(config.App.PublicPath, "static"))
	router.Static("/src", filepath.Join(config.App.PublicPath, "app", "src"))

	// @BasePath /api
	api := router.Group("/api")

	controller.NewAppController().RegisterRoutes(api)
	controller.NewMenuController().RegisterRoutes(api)
	controller.NewSettingController().RegisterRoutes(api)
	controller.NewUserController().RegisterRoutes(api)

	// API文檔路徑: /swagger/index.html
	if !helper.IsRelease() {
		docs.SwaggerInfo.BasePath = "/api"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler)) // /swagger/index.html
	}

	return router
}
