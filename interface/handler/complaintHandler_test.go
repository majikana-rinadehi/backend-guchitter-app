package handler

import (
	"testing"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/usecase"
	"github.com/backend-guchitter-app/util/errors"
	tu "github.com/backend-guchitter-app/util/testUtils"
	"github.com/backend-guchitter-app/util/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// usecase層のMock
type complaintUsecaseMock struct {
	usecase.ComplaintUseCase
	// usecaseのinterfaceからコピペする
	FakeFindAll              func() ([]*model.Complaint, error)
	FakeFindByAvatarId       func(id int) (*model.Complaint, error)
	FakeCreate               func(complaint model.Complaint) (*model.Complaint, error)
	FakeFindBetweenTimestamp func(from string, to string) ([]*model.Complaint, error)
	FakeDeleteByComplaintId  func(id int) error
}

func (m *complaintUsecaseMock) FindAll() ([]*model.Complaint, error) {
	return m.FakeFindAll()
}
func (m *complaintUsecaseMock) FindByAvatarId(id int) (*model.Complaint, error) {
	return m.FakeFindByAvatarId(id)
}
func (m *complaintUsecaseMock) Create(complaint model.Complaint) (*model.Complaint, error) {
	return m.FakeCreate(complaint)
}
func (m *complaintUsecaseMock) FindBetweenTimestamp(from string, to string) ([]*model.Complaint, error) {
	return m.FakeFindBetweenTimestamp(from, to)
}
func (m *complaintUsecaseMock) DeleteByComplaintId(id int) error {
	return m.FakeDeleteByComplaintId(id)
}

var (
	fakeComplaintList = []*model.Complaint{
		{
			ComplaintId:   1,
			ComplaintText: "あああ",
			AvatarId:      1,
		},
	}
	fakeComplaint = model.Complaint{
		ComplaintId:   1,
		ComplaintText: "あああ",
		AvatarId:      1,
	}
)

func Test_complaintHandler_Index(t *testing.T) {
	mockUsecase := &complaintUsecaseMock{
		FakeFindAll: func() ([]*model.Complaint, error) {
			return fakeComplaintList, nil
		},
	}

	// 異常系のケースで使うmock
	mockUsecaseErr := &complaintUsecaseMock{
		FakeFindAll: func() ([]*model.Complaint, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}

	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}
	tests := []struct {
		name       string
		fields     fields
		wantStatus int
		wantBody   []*model.Complaint
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "Test_complaintHandler_Index_200",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 200,
			wantBody:   fakeComplaintList,
			wantErr:    nil,
		},
		{
			name: "Test_complaintHandler_Index_500",
			fields: fields{
				complaintUseCase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantBody:   nil,
			wantErr:    tu.ErrorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := tu.SetupGinContext()
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
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

func Test_complaintHandler_Search(t *testing.T) {
	mockUsecase := &complaintUsecaseMock{
		FakeFindByAvatarId: func(id int) (*model.Complaint, error) {
			if id == 1 {
				return &fakeComplaint, nil
			}
			return nil, nil
		},
	}

	// 異常系のケースで使うmock
	mockUsecaseErr := &complaintUsecaseMock{
		FakeFindByAvatarId: func(id int) (*model.Complaint, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}

	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}

	tests := []struct {
		name string
		// requestParam
		arg        string
		fields     fields
		wantStatus int
		wantBody   *model.Complaint
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "Test_complaintHandler_Search_200",
			arg:  "1",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 200,
			wantBody:   &fakeComplaint,
			wantErr:    nil,
		},
		{
			name: "Test_complaintHandler_Search_500",
			arg:  "1",
			fields: fields{
				complaintUseCase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantBody:   &fakeComplaint,
			wantErr:    tu.ErrorJson,
		},
		{
			name: "Test_complaintHandler_Search_404",
			arg:  "999",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 404,
			wantBody:   &fakeComplaint,
			wantErr:    tu.NotFoundJson,
		},
		{
			name: "Test_complaintHandler_Search_400_empty",
			arg:  "",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 400,
			wantBody:   nil,
			wantErr: &errors.ErrorStruct{
				Message: "Bad request.",
				Fields: []string{
					"Param 'id' is required.",
				},
			},
		},
		{
			name: "Test_complaintHandler_Search_400_space",
			arg:  " ",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 400,
			wantBody:   nil,
			wantErr: &errors.ErrorStruct{
				Message: "Bad request.",
				Fields: []string{
					"Param 'id' is required.",
				},
			},
		},
		{
			name: "Test_complaintHandler_Search_400_not_number",
			arg:  "あ",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 400,
			wantBody:   nil,
			wantErr: &errors.ErrorStruct{
				Message: "Bad request.",
				Fields: []string{
					"Param 'id' must be a 'number'.",
				},
			},
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
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
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

func Test_complaintHandler_Create(t *testing.T) {
	mockUsecase := &complaintUsecaseMock{
		FakeCreate: func(complaint model.Complaint) (*model.Complaint, error) {
			return &complaint, nil
		},
	}
	mockUsecaseErr := &complaintUsecaseMock{
		FakeCreate: func(complaint model.Complaint) (*model.Complaint, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}

	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}
	tests := []struct {
		name string
		// requestBody
		arg        *model.Complaint
		fields     fields
		wantStatus int
		wantBody   *model.Complaint
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "Test_complaintHandler_Create_200",
			arg:  &fakeComplaint,
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 200,
			wantBody:   &fakeComplaint,
			wantErr:    nil,
		},
		{
			name: "Test_complaintHandler_Create_500",
			arg:  &fakeComplaint,
			fields: fields{
				complaintUseCase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantBody:   &fakeComplaint,
			wantErr:    tu.ErrorJson,
		},
		// TODO: 構造体にフィールドが存在しなくても、
		// *gin.Context.BindJSON()で 初期値が次のように入れられてしまう
		// (golangの`zero value`)
		// ので、フィールドが存在しないエラーのテストは一旦見送り
		// {
		// 	ComplaintId: 0,
		// 	ComplaintText: "",
		// 	AvatarId: 0
		// }
		// {
		// 	name: "Test_complaintHandler_Create_400_nil",
		// 	arg: &model.Complaint{
		// 		ComplaintText: "テスト",
		// 	},
		// 	fields: fields{
		// 		complaintUseCase: mockUsecase,
		// 	},
		// 	wantStatus: 400,
		// 	wantBody:   &fakeComplaint,
		// 	wantErr: &errors.ErrorStruct{
		// 		Message: "Bad request.",
		// 		Fields: []string{
		// 			"Param 'complaintId' is required.",
		// 			"Param 'avatarId' is required.",
		// 		},
		// 	},
		// },
		{
			name: "Test_complaintHandler_Create_400_ComplaintText_empty",
			arg: &model.Complaint{
				ComplaintId:   1,
				AvatarId:      1,
				ComplaintText: "",
			},
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 400,
			wantBody:   &fakeComplaint,
			wantErr: &errors.ErrorStruct{
				Message: "Bad request.",
				Fields: utils.SortStrings(
					[]string{
						"Param 'complaintText' is required.",
					},
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w, c := tu.SetupGinContext(tu.WithBody(tt.arg))
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
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

func Test_complaintHandler_FindBetweenTimestamp(t *testing.T) {
	mockUsecase := &complaintUsecaseMock{
		FakeFindBetweenTimestamp: func(from, to string) ([]*model.Complaint, error) {
			if from != "" {
				return fakeComplaintList, nil
			}
			return []*model.Complaint{}, nil
		},
	}
	mockUsecaseErr := &complaintUsecaseMock{
		FakeFindBetweenTimestamp: func(from, to string) ([]*model.Complaint, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}

	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}

	type args struct {
		from string
		to   string
	}

	tests := []struct {
		name string
		args args
		// requestBody
		fields     fields
		wantStatus int
		wantBody   []*model.Complaint
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "FindBetweenTimestamp_200",
			args: args{
				from: "2022-01-01",
				to:   "2022-01-01",
			},
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 200,
			wantBody:   fakeComplaintList,
			wantErr:    nil,
		},
		{
			name: "FindBetweenTimestamp_404",
			args: args{
				from: "",
				to:   "",
			},
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 404,
			wantBody:   fakeComplaintList,
			wantErr:    tu.NotFoundJson,
		},
		{
			name: "FindBetweenTimestamp_500",
			args: args{
				from: "2022-01-01",
				to:   "2022-01-01",
			},
			fields: fields{
				complaintUseCase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantBody:   fakeComplaintList,
			wantErr:    tu.ErrorJson,
		},
		{
			name: "FindBetweenTimestamp_400_not_date_1",
			args: args{
				from: "2022-01-011",
				to:   "2022-01-011",
			},
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 400,
			wantBody:   nil,
			wantErr: &errors.ErrorStruct{
				Message: "Bad request.",
				Fields: []string{
					"Param 'from' must be a 'YYYY-MM-DD'.",
					"Param 'to' must be a 'YYYY-MM-DD'.",
				},
			},
		},
		{
			name: "FindBetweenTimestamp_400_not_date_2",
			args: args{
				from: "2022-01-32",
				to:   "2022-01-32",
			},
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 400,
			wantBody:   nil,
			wantErr: &errors.ErrorStruct{
				Message: "Bad request.",
				Fields: []string{
					"Param 'from' must be a 'YYYY-MM-DD'.",
					"Param 'to' must be a 'YYYY-MM-DD'.",
				},
			},
		},
		{
			name: "FindBetweenTimestamp_400_not_date_3_leap",
			args: args{
				from: "2024-02-29", // OK
				to:   "2024-02-30", // NG
			},
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 400,
			wantBody:   nil,
			wantErr: &errors.ErrorStruct{
				Message: "Bad request.",
				Fields: []string{
					// "Param 'from' must be a 'YYYY-MM-DD'.",
					"Param 'to' must be a 'YYYY-MM-DD'.",
				},
			},
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
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
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

func Test_complaintHandler_DeleteByComplaintId(t *testing.T) {
	mockUsecase := &complaintUsecaseMock{
		FakeDeleteByComplaintId: func(id int) error {
			return nil
		},
	}
	mockUsecaseErr := &complaintUsecaseMock{
		FakeDeleteByComplaintId: func(id int) error {
			return gorm.ErrRecordNotFound
		},
	}
	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}
	tests := []struct {
		name string
		// requestBody
		fields     fields
		wantStatus int
		wantErr    *errors.ErrorStruct
	}{
		// TODO: Add test cases.
		{
			name: "DeleteByComplaintId_204",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 204,
			wantErr:    nil,
		},
		{
			name: "DeleteByComplaintId_500",
			fields: fields{
				complaintUseCase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantErr:    tu.ErrorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, c := tu.SetupGinContext()
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.DeleteByComplaintId(c)

			tu.AssertStatusCode(t, c, tt.wantStatus)
		})
	}
}
