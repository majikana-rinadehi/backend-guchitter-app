package usecase

import (
	"reflect"
	"testing"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/domain/repository"
	"gorm.io/gorm"
)

// Repository層のMock
type avatarRepoMock struct {
	repository.AvatarRepository
	FakeFindAll              func() ([]*model.Avatar, error)
	FakeFindByAvatarId       func(id int) (*model.Avatar, error)
	FakeCreate               func(avatar model.Avatar) (*model.Avatar, error)
	FakeFindBetweenTimestamp func(from string, to string) ([]*model.Avatar, error)
	FakeDeleteByAvatarId     func(id int) error
}

func (m *avatarRepoMock) FindAll() ([]*model.Avatar, error) {
	return m.FakeFindAll()
}

func (m *avatarRepoMock) FindByAvatarId(id int) (*model.Avatar, error) {
	return m.FakeFindByAvatarId(id)
}

func (m *avatarRepoMock) Create(avatar model.Avatar) (*model.Avatar, error) {
	return m.FakeCreate(avatar)
}

func (m *avatarRepoMock) FindBetweenTimestamp(from string, to string) ([]*model.Avatar, error) {
	return m.FakeFindBetweenTimestamp(from, to)
}

func (m *avatarRepoMock) DeleteByAvatarId(id int) error {
	return m.FakeDeleteByAvatarId(id)
}

var (
	// ダミーのavatar
	fakeAvatar = &model.Avatar{
		AvatarId:   1234567890,
		AvatarName: "Nino",
		AvatarText: "なのよ",
		ImageUrl:   "https://hoge.com/fuga",
		Color:      "#f6f6f6",
	}

	fakeAvatarList = []*model.Avatar{
		fakeAvatar,
	}
)

func Test_avatarUseCase_FindAll(t *testing.T) {
	mockRepo := avatarRepoMock{
		FakeFindAll: func() ([]*model.Avatar, error) {
			return fakeAvatarList, nil
		},
	}

	type fields struct {
		avatarRepository repository.AvatarRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.Avatar
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarUseCase_FindAll_normal",
			fields: fields{
				avatarRepository: &mockRepo,
			},
			want:    fakeAvatarList,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := avatarUseCase{
				avatarRepository: tt.fields.avatarRepository,
			}
			got, _ := cu.FindAll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avatarUseCase.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_avatarUseCase_FindByAvatarId(t *testing.T) {
	mockRepo := avatarRepoMock{
		FakeFindByAvatarId: func(id int) (*model.Avatar, error) {
			if id == 999 {
				return nil, gorm.ErrRecordNotFound
			}
			return fakeAvatar, nil
		},
	}

	type fields struct {
		avatarRepository repository.AvatarRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Avatar
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarUseCase_FindByAvatarId_normal",
			fields: fields{
				avatarRepository: &mockRepo,
			},
			args: args{
				id: 1,
			},
			want:    fakeAvatar,
			wantErr: nil,
		},
		{
			name: "Test_avatarUseCase_FindByAvatarId_anormal",
			fields: fields{
				avatarRepository: &mockRepo,
			},
			args: args{
				id: 999,
			},
			want:    nil,
			wantErr: gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := avatarUseCase{
				avatarRepository: tt.fields.avatarRepository,
			}
			got, err := cu.FindByAvatarId(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("avatarUseCase.FindByAvatarId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avatarUseCase.FindByAvatarId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_avatarUseCase_Create(t *testing.T) {
	mockRepo := avatarRepoMock{
		FakeCreate: func(avatar model.Avatar) (*model.Avatar, error) {
			return fakeAvatar, nil
		},
	}

	type fields struct {
		avatarRepository repository.AvatarRepository
	}
	type args struct {
		avatar model.Avatar
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Avatar
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarUseCase_Create_normal",
			fields: fields{
				avatarRepository: &mockRepo,
			},
			args: args{
				avatar: *fakeAvatar,
			},
			want:    fakeAvatar,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := avatarUseCase{
				avatarRepository: tt.fields.avatarRepository,
			}
			got, err := cu.Create(tt.args.avatar)
			if err != tt.wantErr {
				t.Errorf("avatarUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avatarUseCase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_avatarUseCase_FindBetweenTimestamp(t *testing.T) {

	mockRepo := avatarRepoMock{
		FakeFindBetweenTimestamp: func(from, to string) ([]*model.Avatar, error) {
			// TODO write date check
			return fakeAvatarList, nil
		},
	}

	type fields struct {
		avatarRepository repository.AvatarRepository
	}
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Avatar
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarUseCase_FindBetweenTimestamp_normal",
			fields: fields{
				avatarRepository: &mockRepo,
			},
			args: args{
				from: "2022-01-01 12:00:00.000",
				to:   "2023-01-01 12:00:00.000",
			},
			want:    fakeAvatarList,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := avatarUseCase{
				avatarRepository: tt.fields.avatarRepository,
			}
			got, _ := cu.FindBetweenTimestamp(tt.args.from, tt.args.to)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avatarUseCase.FindBetweenTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_avatarUseCase_DeleteByAvatarId(t *testing.T) {
	mockRepo := avatarRepoMock{
		FakeDeleteByAvatarId: func(id int) error {
			if id == 999 {
				return gorm.ErrRecordNotFound
			}
			return nil
		},
	}

	type fields struct {
		avatarRepository repository.AvatarRepository
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
			name: "Test_avatarUseCase_DeleteByAvatarId_normal",
			fields: fields{
				avatarRepository: &mockRepo,
			},
			args: args{
				id: 1,
			},
			wantErr: nil,
		},
		{
			name: "Test_avatarUseCase_DeleteByAvatarId_anormal",
			fields: fields{
				avatarRepository: &mockRepo,
			},
			args: args{
				id: 999,
			},
			wantErr: gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cu := avatarUseCase{
				avatarRepository: tt.fields.avatarRepository,
			}
			if err := cu.DeleteByAvatarId(tt.args.id); err != tt.wantErr {
				t.Errorf("avatarUseCase.DeleteByAvatarId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
