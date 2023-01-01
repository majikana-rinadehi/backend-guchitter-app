package repository

import "github.com/backend-guchitter-app/domain/model"

type AvatarRepository interface {
	FindAll() ([]*model.Avatar, error)
	FindByAvatarId(id int) (*model.Avatar, error)
	Create(avatar model.Avatar) (*model.Avatar, error)
	FindBetweenTimestamp(from string, to string) ([]*model.Avatar, error)
	DeleteByAvatarId(id int) error
}
