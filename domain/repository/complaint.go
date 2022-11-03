package repository

import "example.com/main/domain/model"

type ComplaintRepository interface {
	FindAll() ([]*model.Complaint, error)
	FindByAvatarId(id int) (*model.Complaint, error)
	Create(complaint model.Complaint) (*model.Complaint, error)
}
