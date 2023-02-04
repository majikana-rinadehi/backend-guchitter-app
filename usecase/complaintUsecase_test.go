package usecase

import (
	"reflect"
	"testing"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/domain/repository"
	"gorm.io/gorm"
)

// Repository層のMock
type complaintRepoMock struct {
	repository.ComplaintRepository
	// Repositoryのinterfaceからコピペする
	FakeFindAll        func() ([]*model.Complaint, error)
	FakeFindByAvatarId func(id int) (*model.Complaint, error)
}

func (m *complaintRepoMock) FindAll() ([]*model.Complaint, error) {
	return m.FakeFindAll()
}

func (m *complaintRepoMock) FindByAvatarId(id int) (*model.Complaint, error) {
	return m.FakeFindByAvatarId(id)
}

func Test_complaintUseCase_FindAll(t *testing.T) {
	mockRepo := &complaintRepoMock{
		// mock関数の実装
		FakeFindAll: func() ([]*model.Complaint, error) {
			complaintList := []*model.Complaint{
				{
					ComplaintId:   1,
					ComplaintText: "あああ",
					AvatarId:      1,
				},
			}
			return complaintList, nil
		},
	}
	type fields struct {
		complaintRepository repository.ComplaintRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.Complaint
		wantErr error
	}{
		// TODO: Add test cases.
		{"fetchAll",
			fields{complaintRepository: mockRepo},
			[]*model.Complaint{
				{
					ComplaintId:   1,
					ComplaintText: "あああ",
					AvatarId:      1,
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := complaintUseCase{
				complaintRepository: tt.fields.complaintRepository,
			}
			got, err := cu.FindAll()
			if err != nil && err != tt.wantErr {
				t.Errorf("complaintUseCase.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("complaintUseCase.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complaintUseCase_FindByAvatarId(t *testing.T) {
	mockRepo := &complaintRepoMock{
		FakeFindByAvatarId: func(id int) (*model.Complaint, error) {
			var complaint *model.Complaint = nil
			var err error = nil
			switch id {
			case 1:
				complaint = &model.Complaint{
					ComplaintId:   1,
					AvatarId:      1,
					ComplaintText: "uuu",
				}
			default:
				err = gorm.ErrRecordNotFound
			}
			return complaint, err
		},
	}

	type fields struct {
		complaintRepository repository.ComplaintRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Complaint
		wantErr error
	}{
		// TODO: Add test cases.
		// 正常系
		{
			name: "findByAvatarId_success",
			fields: fields{
				complaintRepository: mockRepo,
				// don't do this causing runtime error 「nil pointer...」
				// complaintRepository: mockRepo.ComplaintRepository,
			},
			args: args{id: 1},
			want: &model.Complaint{
				ComplaintId:   1,
				AvatarId:      1,
				ComplaintText: "uuu",
			},
			wantErr: nil,
		},
		// 異常系
		// // レコードなし
		{
			name: "findByAvatarId_fail_notfound",
			fields: fields{
				complaintRepository: mockRepo,
			},
			args:    args{id: 2},
			want:    nil,
			wantErr: gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := complaintUseCase{
				complaintRepository: tt.fields.complaintRepository,
			}
			got, err := cu.FindByAvatarId(tt.args.id)
			if err != nil && err != tt.wantErr {
				t.Errorf("complaintUseCase.FindByAvatarId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("complaintUseCase.FindByAvatarId() = %v, want %v", got, tt.want)
			}
		})
	}
}
