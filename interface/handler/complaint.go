package handler

import (
	"net/http"

	"example.com/main/usecase"
	"github.com/gin-gonic/gin"
)

type ComplaintHandler interface {
	Index(c *gin.Context)
	Search(c *gin.Context)
}

type complaintHandler struct {
	complaintUseCase usecase.ComplaintUseCase
}

func NewComplaintHandler(cu usecase.ComplaintUseCase) ComplaintHandler {
	return &complaintHandler{
		complaintUseCase: cu,
	}
}

func (ch complaintHandler) Index(c *gin.Context) {
	complaints, err := ch.complaintUseCase.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.IndentedJSON(http.StatusOK, complaints)
}

func (ch complaintHandler) Search(c *gin.Context) {
	id := c.Param("id")
	complaint, err := ch.complaintUseCase.FindByAvatarId(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.IndentedJSON(http.StatusOK, complaint)
}
