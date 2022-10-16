package repository

import "example.com/main/domain/model"

type ComplaintRepository interface {
	FindAll() ([]*model.Complaint, error)
	FindByAvatarId(id string) (*model.Complaint, error)
}
