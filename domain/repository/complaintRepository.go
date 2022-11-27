package repository

import "github.com/backend-guchitter-app/domain/model"

type ComplaintRepository interface {
	FindAll() ([]*model.Complaint, error)
	FindByAvatarId(id int) (*model.Complaint, error)
	Create(complaint model.Complaint) (*model.Complaint, error)
	FindBetweenTimestamp(from string, to string) ([]*model.Complaint, error)
}
