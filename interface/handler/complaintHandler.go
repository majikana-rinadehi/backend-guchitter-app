package handler

import (
	"net/http"

	"strconv"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/logging"
	"github.com/backend-guchitter-app/usecase"
	"github.com/backend-guchitter-app/util/errors"
	"github.com/bloom42/rz-go"
	"github.com/gin-gonic/gin"
)

type ComplaintHandler interface {
	// Index is the handler to fetch all complaints.
	Index(c *gin.Context)
	Search(c *gin.Context)
	Create(c *gin.Context)
	FindBetweenTimestamp(c *gin.Context)
	DeleteByComplaintId(c *gin.Context)
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
		logging.Log.Error("Failed at FindAll()", rz.Err(err))
		// c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		// c.JSON(http.StatusInternalServerError, &errors.ErrorStruct{
		// 	Message: "Internal Server Error",
		// })
		c.JSON(http.StatusInternalServerError, &errors.ErrorStruct{
			Message: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, complaints)
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
	// 必須チェック
	// validateErr := errors.Validate(c.Param("id"))
	// fmt.Println(validateErr)
	// if validateErr != nil {
	// 	for _, err := range validateErr.(validator.ValidationErrors) {
	// 		fieldName, typ := err.Field(), err.Tag()
	// 		fmt.Println(fieldName, typ)
	// 		switch typ {
	// 		case "custom_required":
	// 			c.JSON(http.StatusBadRequest, errors.ErrorStruct{
	// 				Messages: ,
	// 			}})
	// 			return
	// 		case "number":
	// 			c.JSON(http.StatusBadRequest, gin.H{"message": errors.InvalidTypeErrMsg("id", "number")})
	// 			return
	// 		default:

	// 		}
	// 	}
	// }

	if errMessages := errors.Validate(c.Param("id"), "id"); len(errMessages) > 0 {
		c.JSON(http.StatusBadRequest, errors.ErrorStruct{
			Message: "Bad request.",
			Fields:  errMessages,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	complaint, err := ch.complaintUseCase.FindByAvatarId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	if complaint == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, complaint)
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// FindBetweenTimestamp
// @Summary 更新日時がfrom, toの間のComplaintを返す
// @Tags Complaints
// @Produce json
// @Param from query string false "2022-11-27 0:00:00"
// @Param to query string false "2022-11-28 0:00:00"
// @Success 200 {array} model.Complaint
// @Failure 400
// @Failure 500
// @Router /complaints/between-time [get]
func (ch complaintHandler) FindBetweenTimestamp(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	complaintList, err := ch.complaintUseCase.FindBetweenTimestamp(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	if len(complaintList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, complaintList)
}

// DeleteByComplaintId
// @Summary complaintIdで指定したComplaintを1件削除する
// @Tags Complaints
// @Produce json
// @Param id path string false "愚痴ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /complaints/{id} [delete]
func (ch complaintHandler) DeleteByComplaintId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ch.complaintUseCase.DeleteByComplaintId(id)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	// unittestだと200になってしまう、なぜ？？？
	c.Status(http.StatusNoContent)
}
