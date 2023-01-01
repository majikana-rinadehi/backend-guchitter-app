package handler

import (
	"net/http"

	"strconv"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/logging"
	"github.com/backend-guchitter-app/usecase"
	"github.com/bloom42/rz-go"
	"github.com/gin-gonic/gin"
)

type AvatarHandler interface {
	// Index is the handler to fetch all avatars.
	Index(c *gin.Context)
	Search(c *gin.Context)
	Create(c *gin.Context)
	FindBetweenTimestamp(c *gin.Context)
	DeleteByAvatarId(c *gin.Context)
}

type avatarHandler struct {
	avatarUseCase usecase.AvatarUseCase
}

// NewAvatarHandler is the initializer.
func NewAvatarHandler(cu usecase.AvatarUseCase) AvatarHandler {
	return &avatarHandler{
		avatarUseCase: cu,
	}
}

// Index
// @Summary Avatarsを全件取得
// @Tags Avatars
// @Produce json
// @Success 200 {array} model.Avatar
// @Failure 400
// @Failure 500
// @Router /avatars [get]
func (ch avatarHandler) Index(c *gin.Context) {
	avatars, err := ch.avatarUseCase.FindAll()
	if err != nil {
		logging.Log.Error("Failed at FindAll()", rz.Err(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.IndentedJSON(http.StatusOK, avatars)
}

// Search
// @Summary avatarIdで検索したAvatarを1件返す
// @Tags Avatars
// @Produce json
// @Param id path string false "アバターID"
// @Success 200 {object} model.Avatar
// @Failure 400
// @Failure 500
// @Router /avatars/{id} [get]
func (ch avatarHandler) Search(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	avatar, err := ch.avatarUseCase.FindByAvatarId(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	if avatar == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}
	c.IndentedJSON(http.StatusOK, avatar)
}

// Create
// @Summary Avatarを一件登録する
// @Tags Avatars
// @Produce json
// @Param body body model.Avatar false "Avatar"
// @Success 200 {object} model.Avatar "登録したAvatar"
// @Failure 400
// @Failure 500
// @Router /avatars [post]
func (ch avatarHandler) Create(c *gin.Context) {
	var newAvatar *model.Avatar

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAvatar); err != nil {
		return
	}
	result, err := ch.avatarUseCase.Create(*newAvatar)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.IndentedJSON(http.StatusOK, result)
}

// FindBetweenTimestamp
// @Summary 更新日時がfrom, toの間のAvatarを返す
// @Tags Avatars
// @Produce json
// @Param from query string false "2022-11-27 0:00:00"
// @Param to query string false "2022-11-28 0:00:00"
// @Success 200 {array} model.Avatar
// @Failure 400
// @Failure 500
// @Router /avatars/between-time [get]
func (ch avatarHandler) FindBetweenTimestamp(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	avatarList, err := ch.avatarUseCase.FindBetweenTimestamp(from, to)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	if len(avatarList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}
	c.IndentedJSON(http.StatusOK, avatarList)
}

// DeleteByAvatarId
// @Summary avatarIdで指定したAvatarを1件削除する
// @Tags Avatars
// @Produce json
// @Param id path string false "アバターID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /avatars/{id} [delete]
func (ch avatarHandler) DeleteByAvatarId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ch.avatarUseCase.DeleteByAvatarId(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.Status(http.StatusNoContent)
}
