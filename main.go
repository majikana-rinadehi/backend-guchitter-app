package main

import (
	"fmt"

	"example.com/main/infrastructure/persistence"
	"example.com/main/interface/handler"
	"example.com/main/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	complaintPersistence := persistence.NewComplaintPersistence()
	complaintUseCase := usecase.NewComplaintUseCase(complaintPersistence)
	complaintHandler := handler.NewComplaintHandler(complaintUseCase)

	router := gin.Default()
	router.GET("/complaints", complaintHandler.Index)
	router.GET("/complaints/:id", complaintHandler.Search)
	router.POST("/complaints", complaintHandler.Create)

	fmt.Println("sa-bakidou")
	router.Run("localhost:8080")
}
