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
	FakeFindAll              func() ([]*model.Complaint, error)
	FakeFindByAvatarId       func(id int) (*model.Complaint, error)
	FakeCreate               func(complaint model.Complaint) (*model.Complaint, error)
	FakeFindBetweenTimestamp func(from string, to string) ([]*model.Complaint, error)
	FakeDeleteByComplaintId  func(id int) error
}

func (m *complaintRepoMock) FindAll() ([]*model.Complaint, error) {
	return m.FakeFindAll()
}

func (m *complaintRepoMock) FindByAvatarId(id int) (*model.Complaint, error) {
	return m.FakeFindByAvatarId(id)
}

func (m *complaintRepoMock) Create(complaint model.Complaint) (*model.Complaint, error) {
	return m.FakeCreate(complaint)
}

func (m *complaintRepoMock) FindBetweenTimestamp(from, to string) ([]*model.Complaint, error) {
	return m.FakeFindBetweenTimestamp(from, to)
}

func (m *complaintRepoMock) DeleteByComplaintId(id int) error {
	return m.FakeDeleteByComplaintId(id)
}

var (
	// ダミーのComplaint
	fakeComplaint = &model.Complaint{
		ComplaintId:   1,
		AvatarId:      1,
		ComplaintText: "uuu",
	}
)

func Test_complaintUseCase_FindAll(t *testing.T) {
	mockRepo := &complaintRepoMock{
		// mock関数の実装
		FakeFindAll: func() ([]*model.Complaint, error) {
			complaintList := []*model.Complaint{
				fakeComplaint,
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
				fakeComplaint,
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
				complaint = fakeComplaint
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
			args:    args{id: 1},
			want:    fakeComplaint,
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

func Test_complaintUseCase_Create(t *testing.T) {
	mockRepo := complaintRepoMock{
		FakeCreate: func(complaint model.Complaint) (*model.Complaint, error) {
			return &complaint, nil
		},
	}

	type fields struct {
		complaintRepository repository.ComplaintRepository
	}
	type args struct {
		complaint model.Complaint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Complaint
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "Test_complaintUseCase_Create",
			fields: fields{
				complaintRepository: &mockRepo,
			},
			args: args{
				model.Complaint{
					ComplaintId:   1,
					ComplaintText: "あああ",
					AvatarId:      1,
				},
			},
			want: &model.Complaint{
				ComplaintId:   1,
				ComplaintText: "あああ",
				AvatarId:      1,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := complaintUseCase{
				complaintRepository: tt.fields.complaintRepository,
			}
			got, _ := cu.Create(tt.args.complaint)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("complaintUseCase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complaintUseCase_FindBetweenTimestamp(t *testing.T) {
	mockRepo := complaintRepoMock{
		FakeFindBetweenTimestamp: func(from, to string) ([]*model.Complaint, error) {
			return []*model.Complaint{
				fakeComplaint,
			}, nil
		},
	}
	type fields struct {
		complaintRepository repository.ComplaintRepository
	}
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Complaint
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "Test_complaintUseCase_FindBetweenTimestamp",
			fields: fields{
				complaintRepository: &mockRepo,
			},
			args: args{
				from: "2022-02-01",
				to:   "2022-02-28",
			},
			want: []*model.Complaint{
				fakeComplaint,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := complaintUseCase{
				complaintRepository: tt.fields.complaintRepository,
			}
			got, _ := cu.FindBetweenTimestamp(tt.args.from, tt.args.to)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("complaintUseCase.FindBetweenTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complaintUseCase_DeleteByComplaintId(t *testing.T) {
	mockRepo := &complaintRepoMock{
		FakeDeleteByComplaintId: func(id int) error {
			if id == 999 {
				return gorm.ErrRecordNotFound
			}
			return nil
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
		wantErr error
	}{
		// TODO: Add test cases.
		{
			// 正常系
			name: "Test_complaintUseCase_DeleteByComplaintId_normal",
			fields: fields{
				complaintRepository: mockRepo,
			},
			args: args{
				id: 1,
			},
			wantErr: nil,
		},
		{
			// 異常系
			name: "Test_complaintUseCase_DeleteByComplaintId_anormal",
			fields: fields{
				complaintRepository: mockRepo,
			},
			args: args{
				id: 999,
			},
			wantErr: gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := complaintUseCase{
				complaintRepository: tt.fields.complaintRepository,
			}
			if err := cu.DeleteByComplaintId(tt.args.id); err != tt.wantErr {
				t.Errorf("complaintUseCase.DeleteByComplaintId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
