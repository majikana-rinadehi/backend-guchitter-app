package main

import (
	"example.com/main/config"
	_ "example.com/main/docs"
	"example.com/main/infrastructure/persistence"
	"example.com/main/interface/handler"
	logging "example.com/main/logging"
	"example.com/main/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

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
	corsConf.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(corsConf))

	router.GET("/complaints", complaintHandler.Index)
	router.GET("/complaints/:id", complaintHandler.Search)
	router.POST("/complaints", complaintHandler.Create)

	// http://localhost:8080/swagger/index.html にswagger UI を表示する
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8080")
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
