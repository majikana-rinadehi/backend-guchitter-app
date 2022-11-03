package handler

import (
	"net/http"

	"strconv"

	"example.com/main/domain/model"
	"example.com/main/usecase"
	"github.com/gin-gonic/gin"
)

type ComplaintHandler interface {
	// Index is the handler to fetch all complaints.
	Index(c *gin.Context)
	Search(c *gin.Context)
	Create(c *gin.Context)
}

type complaintHandler struct {
	complaintUseCase usecase.ComplaintUseCase
}

// NewComplaintHandler is the initializer.
func NewComplaintHandler(cu usecase.ComplaintUseCase) ComplaintHandler {
	return &complaintHandler{
		complaintUseCase: cu,
	}
}

// Index
// @Summary Complaintsを全件取得
// @Tags Complaints
// @Produce json
// @Success 200 {array} model.Complaint
// @Failure 400
// @Failure 500
// @Router /complaints [get]
func (ch complaintHandler) Index(c *gin.Context) {
	complaints, err := ch.complaintUseCase.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.IndentedJSON(http.StatusOK, complaints)
}

// Search
// @Summary avatarIdで検索したComplaintを1件返す
// @Tags Complaints
// @Produce json
// @Param id path string false "アバターID"
// @Success 200 {object} model.Complaint
// @Failure 400
// @Failure 500
// @Router /complaints/{id} [get]
func (ch complaintHandler) Search(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	complaint, err := ch.complaintUseCase.FindByAvatarId(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	if complaint == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}
	c.IndentedJSON(http.StatusOK, complaint)
}

// Create
// @Summary Complaintを一件登録する
// @Tags Complaints
// @Produce json
// @Param body body model.Complaint false "Complaint"
// @Success 200 {object} model.Complaint "登録したComplaint"
// @Failure 400
// @Failure 500
// @Router /complaints [post]
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
