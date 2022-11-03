package usecase

import (
	"example.com/main/domain/model"
	"example.com/main/domain/repository"
)

type ComplaintUseCase interface {
	FindAll() ([]*model.Complaint, error)
	FindByAvatarId(id int) (*model.Complaint, error)
	Create(complaint model.Complaint) (*model.Complaint, error)
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

func (cu complaintUseCase) FindByAvatarId(id int) (*model.Complaint, error) {
	complaint, err := cu.complaintRepository.FindByAvatarId(id)
	return complaint, err
}

func (cu complaintUseCase) Create(complaint model.Complaint) (*model.Complaint, error) {
	result, err := cu.complaintRepository.Create(complaint)
	return result, err
}
