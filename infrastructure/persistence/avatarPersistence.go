package persistence

import (
	"errors"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/domain/repository"
	"github.com/backend-guchitter-app/logging"
	"github.com/bloom42/rz-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type avatarPersistence struct {
	Conn *gorm.DB
}

func NewAvatarPersistence(conn *gorm.DB) repository.AvatarRepository {
	return &avatarPersistence{
		Conn: conn,
	}
}

func (cp *avatarPersistence) FindAll() (avatarList []*model.Avatar, err error) {
	db := cp.Conn

	if err := db.Find(&avatarList).Error; err != nil {
		return nil, err
	}

	return avatarList, nil
}

func (cp *avatarPersistence) FindByAvatarId(id int) (avatar *model.Avatar, err error) {
	db := cp.Conn

	// Typeormみたくカラム名をキャメルケース(avatarId)にするとエラー
	err = db.First(&avatar, "avatar_id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		avatar = nil
	} else if err != nil {
		return nil, err
	}

	return avatar, nil
}

func (cp *avatarPersistence) Create(avatar model.Avatar) (*model.Avatar, error) {
	db := cp.Conn

	result := db.Create(&avatar)

	if result.Error != nil {
		return nil, result.Error
	}

	logging.Log.Debug("avatar", rz.Any("avatar", avatar))

	return &avatar, nil
}

func (cp *avatarPersistence) FindBetweenTimestamp(from, to string) (avatarList []*model.Avatar, err error) {
	db := cp.Conn

	chain := db.Where("")
	if from != "" {
		chain.Where("last_update >= ?", from)
	}
	if to != "" {
		chain.Where("last_update <= ?", to)
	}

	if err := chain.Find(&avatarList).Error; err != nil {
		return nil, err
	}

	return avatarList, nil
}

func (cp *avatarPersistence) DeleteByAvatarId(id int) error {
	db := cp.Conn

	if err := db.
		Clauses(clause.Returning{}).
		Where("avatar_id = ?", id).
		Delete(&model.Avatar{}).Error; err != nil {
		return err
	}

	return nil
}
