package usecase

import (
	"example.com/main/domain/model"
	"example.com/main/domain/repository"
)

type ComplaintUseCase interface {
	FindAll() ([]*model.Complaint, error)
	FindByAvatarId(id string) (*model.Complaint, error)
}

type complaintUseCase struct {
	complaintRepository repository.ComplaintRepository
}

func NewComplaintUseCase(cr repository.ComplaintRepository) ComplaintUseCase {
	return &complaintUseCase{
		complaintRepository: cr,
	}
}

func (cu complaintUseCase) FindAll() ([]*model.Complaint, error) {
	complaintList, err := cu.complaintRepository.FindAll()
	return complaintList, err
}

func (cu complaintUseCase) FindByAvatarId(id string) (*model.Complaint, error) {
	complaint, err := cu.complaintRepository.FindByAvatarId(id)
	return complaint, err
}
