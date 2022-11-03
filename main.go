package main

import (
	"example.com/main/config"
	_ "example.com/main/docs"
	"example.com/main/infrastructure/persistence"
	"example.com/main/interface/handler"
	"example.com/main/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title gin-swagger guchitter
// @version 0.0.1
// @lisence.name rudy
// @description はじめてのswagger
func main() {
	complaintPersistence := persistence.NewComplaintPersistence(config.Connect())
	complaintUseCase := usecase.NewComplaintUseCase(complaintPersistence)
	complaintHandler := handler.NewComplaintHandler(complaintUseCase)

	router := gin.Default()
	router.GET("/complaints", complaintHandler.Index)
	router.GET("/complaints/:id", complaintHandler.Search)
	router.POST("/complaints", complaintHandler.Create)

	// http://localhost:8080/swagger/index.html にswagger UI を表示する
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8080")
}
