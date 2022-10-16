package persistence

import (
	"example.com/main/domain/model"
	"example.com/main/domain/repository"
)

type complaintPersistence struct {
}

var complaintList = []*model.Complaint{
	{ComplaintText: "textA", AvatarId: "avatarA"},
	{ComplaintText: "textB", AvatarId: "avatarB"},
	{ComplaintText: "textC", AvatarId: "avatarC"},
}

func NewComplaintPersistence() repository.ComplaintRepository {
	return &complaintPersistence{}
}

func (cp *complaintPersistence) FindAll() ([]*model.Complaint, error) {

	return complaintList, nil
}

func (cp *complaintPersistence) FindByAvatarId(id string) (*model.Complaint, error) {
	var complaint *model.Complaint

	for _, v := range complaintList {
		if v.AvatarId == id {
			complaint = v
		}
	}
	return complaint, nil
}
