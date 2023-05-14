package handler

import (
	"testing"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/usecase"
	"github.com/backend-guchitter-app/util/errors"
	tu "github.com/backend-guchitter-app/util/testUtils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type avatarUsecaseMock struct {
	usecase.AvatarUseCase
	FakeFindAll              func() ([]*model.Avatar, error)
	FakeFindByAvatarId       func(id int) (*model.Avatar, error)
	FakeCreate               func(avatar model.Avatar) (*model.Avatar, error)
	FakeFindBetweenTimestamp func(from string, to string) ([]*model.Avatar, error)
	FakeDeleteByAvatarId     func(id int) error
}

func (m *avatarUsecaseMock) FindAll() ([]*model.Avatar, error) {
	return m.FakeFindAll()
}

func (m *avatarUsecaseMock) FindByAvatarId(id int) (*model.Avatar, error) {
	return m.FakeFindByAvatarId(id)
}

func (m *avatarUsecaseMock) Create(avatar model.Avatar) (*model.Avatar, error) {
	return m.FakeCreate(avatar)
}

func (m *avatarUsecaseMock) FindBetweenTimestamp(from, to string) ([]*model.Avatar, error) {
	return m.FakeFindBetweenTimestamp(from, to)
}

func (m *avatarUsecaseMock) DeleteByAvatarId(id int) error {
	return m.FakeDeleteByAvatarId(id)
}

var (
	fakeAvatarList = []*model.Avatar{
		{
			AvatarId:   1,
			AvatarName: "Nino",
			AvatarText: "なのよ",
			ImageUrl:   "example.com",
			Color:      "#ffffff",
		},
		{
			AvatarId:   2,
			AvatarName: "Miku",
			AvatarText: "なんだよ",
			ImageUrl:   "example.com",
			Color:      "#ffffff",
		},
	}
	fakeAvatar = model.Avatar{
		AvatarId:   1,
		AvatarName: "Nino",
		AvatarText: "なのよ",
		ImageUrl:   "example.com",
		Color:      "#ffffff",
	}
)

func Test_avatarHandler_Index(t *testing.T) {
	mock := &avatarUsecaseMock{
		// TODO: implement mock
		FakeFindAll: func() ([]*model.Avatar, error) {
			return fakeAvatarList, nil
		},
	}

	mockErr := &avatarUsecaseMock{
		// TODO: implement mock for err
		FakeFindAll: func() ([]*model.Avatar, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}
	type fields struct {
		avatarUseCase usecase.AvatarUseCase
	}
	tests := []struct {
		name       string
		fields     fields
		wantStatus int
		wantBody   []*model.Avatar
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarHandler_Index_200",
			fields: fields{
				avatarUseCase: mock,
			},
			wantStatus: 200,
			wantBody:   fakeAvatarList,
			wantErr:    nil,
		},
		{
			name: "Test_avatarHandler_Index_500",
			fields: fields{
				avatarUseCase: mockErr,
			},
			wantStatus: 500,
			wantBody:   nil,
			wantErr:    tu.ErrorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := tu.SetupGinContext()
			ch := avatarHandler{
				avatarUseCase: tt.fields.avatarUseCase,
			}
			ch.Index(c)

			tu.AssertStatusCode(t, c, tt.wantStatus)
			if tt.wantErr != nil {
				tu.AssertResponse(t, w, tt.wantErr)
			} else {
				tu.AssertResponse(t, w, tt.wantBody)
			}
		})
	}
}

func Test_avatarHandler_Search(t *testing.T) {
	var mock = &avatarUsecaseMock{
		// TODO: implement mock
		FakeFindByAvatarId: func(id int) (*model.Avatar, error) {
			if id == 1 {
				return &fakeAvatar, nil
			}
			return nil, nil
		},
	}
	var mockErr = &avatarUsecaseMock{
		// TODO: implement mock for err
		FakeFindByAvatarId: func(id int) (*model.Avatar, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}
	type fields struct {
		avatarUseCase usecase.AvatarUseCase
	}
	tests := []struct {
		name       string
		fields     fields
		arg        string
		wantStatus int
		wantBody   *model.Avatar
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarHandler_Search_200",
			arg:  "1",
			fields: fields{
				avatarUseCase: mock,
			},
			wantStatus: 200,
			wantBody:   &fakeAvatar,
			wantErr:    nil,
		},
		{
			name: "Test_avatarHandler_Search_500",
			arg:  "1",
			fields: fields{
				avatarUseCase: mockErr,
			},
			wantStatus: 500,
			wantBody:   nil,
			wantErr:    tu.ErrorJson,
		},
		{
			name: "Test_avatarHandler_Search_404",
			arg:  "999",
			fields: fields{
				avatarUseCase: mock,
			},
			wantStatus: 404,
			wantBody:   nil,
			wantErr:    tu.NotFoundJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := tu.SetupGinContext(
				tu.WithParam(
					[]gin.Param{
						{Key: "id", Value: tt.arg},
					},
				),
			)
			ch := avatarHandler{
				avatarUseCase: tt.fields.avatarUseCase,
			}
			ch.Search(c)

			tu.AssertStatusCode(t, c, tt.wantStatus)
			if tt.wantErr != nil {
				tu.AssertResponse(t, w, tt.wantErr)
			} else {
				tu.AssertResponse(t, w, tt.wantBody)
			}
		})
	}
}

func Test_avatarHandler_Create(t *testing.T) {
	var mock = &avatarUsecaseMock{
		// TODO: implement mock
		FakeCreate: func(avatar model.Avatar) (*model.Avatar, error) {
			return &fakeAvatar, nil
		},
	}
	var mockErr = &avatarUsecaseMock{
		// TODO: implement mock for err
		FakeCreate: func(avatar model.Avatar) (*model.Avatar, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}
	type fields struct {
		avatarUseCase usecase.AvatarUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name       string
		fields     fields
		arg        *model.Avatar
		wantStatus int
		wantBody   *model.Avatar
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarHandler_Create_200",
			arg:  &fakeAvatar,
			fields: fields{
				avatarUseCase: mock,
			},
			wantStatus: 200,
			wantBody:   &fakeAvatar,
			wantErr:    nil,
		},
		{
			name: "Test_avatarHandler_Create_500",
			arg:  &fakeAvatar,
			fields: fields{
				avatarUseCase: mockErr,
			},
			wantStatus: 500,
			wantBody:   nil,
			wantErr:    tu.ErrorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := tu.SetupGinContext(tu.WithBody(tt.arg))
			ch := avatarHandler{
				avatarUseCase: tt.fields.avatarUseCase,
			}
			ch.Create(c)

			tu.AssertStatusCode(t, c, tt.wantStatus)
			if tt.wantErr != nil {
				tu.AssertResponse(t, w, tt.wantErr)
			} else {
				tu.AssertResponse(t, w, tt.wantBody)
			}
		})
	}
}

func Test_avatarHandler_FindBetweenTimestamp(t *testing.T) {
	var mock = &avatarUsecaseMock{
		// TODO: implement mock
		FakeFindBetweenTimestamp: func(from, to string) ([]*model.Avatar, error) {
			if from != "" {
				return fakeAvatarList, nil
			}
			return nil, nil
		},
	}
	var mockErr = &avatarUsecaseMock{
		// TODO: implement mock for err
		FakeFindBetweenTimestamp: func(from, to string) ([]*model.Avatar, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}
	type fields struct {
		avatarUseCase usecase.AvatarUseCase
	}
	type args struct {
		from, to string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
		wantBody   []*model.Avatar
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarHandler_FindBetweenTimestamp_200",
			args: args{
				from: "2022-01-01",
				to:   "2022-01-01",
			},
			fields: fields{
				avatarUseCase: mock,
			},
			wantStatus: 200,
			wantBody:   fakeAvatarList,
			wantErr:    nil,
		},
		{
			name: "Test_avatarHandler_FindBetweenTimestamp_404",
			args: args{
				from: "",
				to:   "",
			},
			fields: fields{
				avatarUseCase: mock,
			},
			wantStatus: 404,
			wantBody:   fakeAvatarList,
			wantErr:    tu.NotFoundJson,
		},
		{
			name: "Test_avatarHandler_FindBetweenTimestamp_500",
			args: args{
				from: "2022-01-01",
				to:   "2022-01-01",
			},
			fields: fields{
				avatarUseCase: mockErr,
			},
			wantStatus: 500,
			wantBody:   fakeAvatarList,
			wantErr:    tu.ErrorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := tu.SetupGinContext(
				tu.WithQuery(
					[]gin.Param{
						{Key: "from", Value: tt.args.from},
						{Key: "to", Value: tt.args.to},
					},
				),
			)
			ch := avatarHandler{
				avatarUseCase: tt.fields.avatarUseCase,
			}
			ch.FindBetweenTimestamp(c)

			tu.AssertStatusCode(t, c, tt.wantStatus)
			if tt.wantErr != nil {
				tu.AssertResponse(t, w, tt.wantErr)
			} else {
				tu.AssertResponse(t, w, tt.wantBody)
			}
		})
	}
}

func Test_avatarHandler_DeleteByAvatarId(t *testing.T) {
	var mock = &avatarUsecaseMock{
		// TODO: implement mock
		FakeDeleteByAvatarId: func(id int) error {
			return nil
		},
	}
	var mockErr = &avatarUsecaseMock{
		// TODO: implement mock for err
		FakeDeleteByAvatarId: func(id int) error {
			return gorm.ErrRecordNotFound
		},
	}
	type fields struct {
		avatarUseCase usecase.AvatarUseCase
	}
	tests := []struct {
		name       string
		fields     fields
		arg        string
		wantStatus int
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarHandler_DeleteByAvatarId_204",
			arg:  "1",
			fields: fields{
				avatarUseCase: mock,
			},
			wantStatus: 204,
			wantErr:    nil,
		},
		{
			name: "Test_avatarHandler_DeleteByAvatarId_500",
			arg:  "1",
			fields: fields{
				avatarUseCase: mockErr,
			},
			wantStatus: 500,
			wantErr:    tu.ErrorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := tu.SetupGinContext()
			ch := avatarHandler{
				avatarUseCase: tt.fields.avatarUseCase,
			}
			ch.DeleteByAvatarId(c)

			tu.AssertStatusCode(t, c, tt.wantStatus)
		})
	}
}
