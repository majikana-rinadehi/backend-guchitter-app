package usecase

import (
	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/domain/repository"
)

type AvatarUseCase interface {
	FindAll() ([]*model.Avatar, error)
	FindByAvatarId(id int) (*model.Avatar, error)
	Create(avatar model.Avatar) (*model.Avatar, error)
	FindBetweenTimestamp(from string, to string) ([]*model.Avatar, error)
	DeleteByAvatarId(id int) error
}

type avatarUseCase struct {
	avatarRepository repository.AvatarRepository
}

func NewAvatarUseCase(cr repository.AvatarRepository) AvatarUseCase {
	return &avatarUseCase{
		avatarRepository: cr,
	}
}

func (cu avatarUseCase) FindAll() ([]*model.Avatar, error) {
	avatarList, err := cu.avatarRepository.FindAll()
	return avatarList, err
}

func (cu avatarUseCase) FindByAvatarId(id int) (*model.Avatar, error) {
	avatar, err := cu.avatarRepository.FindByAvatarId(id)
	return avatar, err
}

func (cu avatarUseCase) Create(avatar model.Avatar) (*model.Avatar, error) {
	result, err := cu.avatarRepository.Create(avatar)
	return result, err
}

func (cu avatarUseCase) FindBetweenTimestamp(from string, to string) ([]*model.Avatar, error) {
	avatarList, err := cu.avatarRepository.FindBetweenTimestamp(from, to)
	return avatarList, err
}

func (cu avatarUseCase) DeleteByAvatarId(id int) error {
	err := cu.avatarRepository.DeleteByAvatarId(id)
	return err
}
