package persistence

import (
	"errors"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type complaintPersistence struct {
	Conn *gorm.DB
}

func NewComplaintPersistence(conn *gorm.DB) repository.ComplaintRepository {
	return &complaintPersistence{
		Conn: conn,
	}
}

func (cp *complaintPersistence) FindAll() (complaintList []*model.Complaint, err error) {
	db := cp.Conn

	if err := db.Find(&complaintList).Error; err != nil {
		return nil, err
	}

	return complaintList, nil
}

func (cp *complaintPersistence) FindByAvatarId(id int) (complaint *model.Complaint, err error) {
	db := cp.Conn

	// Typeormみたくカラム名をキャメルケース(avatarId)にするとエラー
	err = db.First(&complaint, "avatar_id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		complaint = nil
	} else if err != nil {
		return nil, err
	}

	return complaint, nil
}

func (cp *complaintPersistence) Create(complaint model.Complaint) (*model.Complaint, error) {
	db := cp.Conn

	if result := db.Create(&complaint); result.Error != nil {
		return nil, result.Error
	}

	return &complaint, nil
}

func (cp *complaintPersistence) FindBetweenTimestamp(from, to string) (complaintList []*model.Complaint, err error) {
	db := cp.Conn

	chain := db.Where("")
	if from != "" {
		chain.Where("last_update >= ?", from)
	}
	if to != "" {
		chain.Where("last_update <= ?", to)
	}

	if err := chain.Find(&complaintList).Error; err != nil {
		return nil, err
	}

	return complaintList, nil
}

func (cp *complaintPersistence) DeleteByComplaintId(id int) error {
	db := cp.Conn

	if err := db.
		Clauses(clause.Returning{}).
		Where("complaint_id = ?", id).
		Delete(&model.Complaint{}).Error; err != nil {
		return err
	}

	return nil
}
