package handler

import (
	"net/http"

	"example.com/main/domain/model"
	"example.com/main/usecase"
	"github.com/gin-gonic/gin"
)

type ComplaintHandler interface {
	Index(c *gin.Context)
	Search(c *gin.Context)
	Create(c *gin.Context)
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

func (ch complaintHandler) Create(c *gin.Context) {
	var newComplaint *model.Complaint

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newComplaint); err != nil {
		return
	}
	result, err := ch.complaintUseCase.Create(*newComplaint)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.IndentedJSON(http.StatusOK, result)
}
