package main

import (
	"os"

	"github.com/backend-guchitter-app/config"
	_ "github.com/backend-guchitter-app/docs"
	"github.com/backend-guchitter-app/infrastructure/persistence"
	"github.com/backend-guchitter-app/interface/handler"
	logging "github.com/backend-guchitter-app/logging"
	"github.com/backend-guchitter-app/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	// header name of unique request id
	XRequestId = "X-Request-ID"
)

// @title gin-swagger guchitter
// @version 0.0.1
// @lisence.name rudy
// @description はじめてのswagger
func main() {
	// 依存性の注入
	complaintPersistence := persistence.NewComplaintPersistence(config.Connect())
	complaintUseCase := usecase.NewComplaintUseCase(complaintPersistence)
	complaintHandler := handler.NewComplaintHandler(complaintUseCase)

	router := gin.Default()

	// リクエストID設定
	router.Use(requestid.New())

	// ロギング設定
	router.Use(loggerSetup())

	// CORS設定
	corsConf := cors.DefaultConfig()
	env := os.Getenv("GUCHITTER_ENV")

	if env == "production" {
		env = "production"
	} else {
		env = "development"
	}

	godotenv.Load(".env." + env)
	FRONT_ORIGIN := os.Getenv("guchitter_FRONT_ORIGIN")
	corsConf.AllowOrigins = []string{FRONT_ORIGIN}
	router.Use(cors.New(corsConf))

	router.GET("/complaints", complaintHandler.Index)
	router.GET("/complaints/:id", complaintHandler.Search)
	router.POST("/complaints", complaintHandler.Create)

	// http://localhost:8080/swagger/index.html にswagger UI を表示する
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	router.Run(":" + port)
}

// Logger設定を行うミドルウェア
func loggerSetup() gin.HandlerFunc {
	return func(c *gin.Context) {
		// logging設定
		requestId := c.Request.Header.Get(XRequestId)
		logging.SetupLogger(requestId)
		c.Next()
	}
}
