package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/backend-guchitter-app/domain/model"
	"github.com/backend-guchitter-app/usecase"
	"github.com/backend-guchitter-app/util/errors"
	testutils "github.com/backend-guchitter-app/util/testUtils"
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
	errorJson    = &errors.ErrorStruct{Message: "Internal Server Error"}
	notFoundJson = &errors.ErrorStruct{Message: "Not Found"}
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
			name: "Index_200",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 200,
			wantBody:   fakeComplaintList,
			wantErr:    nil,
		},
		{
			name: "Index_500",
			fields: fields{
				complaintUseCase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantBody:   nil,
			wantErr:    errorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := testutils.SetupGinContext()

			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.Index(c)

			// ステータスコードのアサーション
			if !reflect.DeepEqual(w.Code, tt.wantStatus) {
				t.Errorf("complaintHandler.Index() = %v, want %v", w.Code, tt.wantStatus)
			}
			// レスポンスJSONのアサーション
			// 異常系の場合
			if tt.wantErr != nil {
				// var errorStruct *struct{ Message string }
				var errorStruct *errors.ErrorStruct
				fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				// ↓this is the string expression returned by `Index()`
				test1 := "{\"message\": \"Internal Server Error\"}null"
				test2 := "{\"message\": \"Internal Server Error\"}"
				json.Unmarshal([]byte(test1), &errorStruct)
				json.Unmarshal([]byte(test2), &errorStruct)
				json.Unmarshal(w.Body.Bytes(), &errorStruct)
				if !reflect.DeepEqual(errorStruct, tt.wantErr) {
					t.Errorf("complaintHandler.Index() = %v, want %v", errorStruct, tt.wantErr)
				}
			} else {
				var complaintList []*model.Complaint
				fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				json.Unmarshal(w.Body.Bytes(), &complaintList)
				if !reflect.DeepEqual(complaintList, tt.wantBody) {
					t.Errorf("complaintHandler.Index() = %v, want %v", complaintList, tt.wantBody)
				}
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
			name: "Search_200",
			arg:  "1",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 200,
			wantBody:   &fakeComplaint,
			wantErr:    nil,
		},
		{
			name: "Search_500",
			arg:  "1",
			fields: fields{
				complaintUseCase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantBody:   &fakeComplaint,
			wantErr:    errorJson,
		},
		{
			name: "Search_404",
			arg:  "999",
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 404,
			wantBody:   &fakeComplaint,
			wantErr:    notFoundJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := testutils.SetupGinContextWithParam(
				[]gin.Param{
					{Key: "id", Value: tt.arg},
				},
			)

			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.Search(c)

			// ステータスコードのアサーション
			if !reflect.DeepEqual(w.Code, tt.wantStatus) {
				t.Errorf("complaintHandler.Search() = %v, want %v", w.Code, tt.wantStatus)
			}

			// 異常系のレスポンスJSONのアサーション
			if tt.wantErr != nil {
				var errorStruct *errors.ErrorStruct
				fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				json.Unmarshal(w.Body.Bytes(), &errorStruct)
				if !reflect.DeepEqual(errorStruct, tt.wantErr) {
					t.Errorf("complaintHandler.Search() = %v, want %v", errorStruct, tt.wantErr)
				}
			} else {
				// レスポンスJSONのアサーション
				// ↓it causes nil pointer panic at !reflect.DeepEqual
				// var complaint model.Complaint
				var complaint *model.Complaint
				fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				json.Unmarshal(w.Body.Bytes(), &complaint)
				if !reflect.DeepEqual(complaint, tt.wantBody) {
					t.Errorf("complaintHandler.Search() = %v, want %v", complaint, tt.wantBody)
				}
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
			name: "Create_200",
			arg:  &fakeComplaint,
			fields: fields{
				complaintUseCase: mockUsecase,
			},
			wantStatus: 200,
			wantBody:   &fakeComplaint,
			wantErr:    nil,
		},
		{
			name: "Create_500",
			arg:  &fakeComplaint,
			fields: fields{
				complaintUseCase: mockUsecaseErr,
			},
			wantStatus: 500,
			wantBody:   &fakeComplaint,
			wantErr:    errorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w, c := testutils.SetupGinContextWithBody(tt.arg)
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.Create(c)

			// ステータスコードのアサーション
			if !reflect.DeepEqual(w.Code, tt.wantStatus) {
				t.Errorf("complaintHandler.Create() = %v, want %v", w.Code, tt.wantStatus)
			}

			// 異常系のレスポンスJSONのアサーション
			if tt.wantErr != nil {
				var errorStruct *errors.ErrorStruct
				fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				json.Unmarshal(w.Body.Bytes(), &errorStruct)
				if !reflect.DeepEqual(errorStruct, tt.wantErr) {
					t.Errorf("complaintHandler.Create() = %v, want %v", errorStruct, tt.wantErr)
				}
			} else {
				// レスポンスJSONのアサーション
				// ↓it causes nil pointer panic at !reflect.DeepEqual
				// var complaint model.Complaint
				var complaint *model.Complaint
				fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				json.Unmarshal(w.Body.Bytes(), &complaint)
				if !reflect.DeepEqual(complaint, tt.wantBody) {
					t.Errorf("complaintHandler.Create() = %v, want %v", complaint, tt.wantBody)
				}
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
			wantErr:    notFoundJson,
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
			wantErr:    errorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := testutils.SetupGinContextWithQuery(
				[]gin.Param{
					{Key: "from", Value: tt.args.from},
					{Key: "to", Value: tt.args.to},
				},
			)
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.FindBetweenTimestamp(c)

			// ステータスコードのアサーション
			if !reflect.DeepEqual(w.Code, tt.wantStatus) {
				t.Errorf("complaintHandler.FindBetweenTimestamp() = %v, want %v", w.Code, tt.wantStatus)
			}

			// 異常系のレスポンスJSONのアサーション
			if tt.wantErr != nil {
				var errorStruct *errors.ErrorStruct
				fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				json.Unmarshal(w.Body.Bytes(), &errorStruct)
				if !reflect.DeepEqual(errorStruct, tt.wantErr) {
					t.Errorf("complaintHandler.FindBetweenTimestamp() = %v, want %v", errorStruct, tt.wantErr)
				}
			} else {
				// レスポンスJSONのアサーション
				// ↓it causes nil pointer panic at !reflect.DeepEqual
				// var complaint model.Complaint
				// var complaintList []*model.Complaint
				// fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				// json.Unmarshal(w.Body.Bytes(), &complaintList)
				// if !reflect.DeepEqual(complaintList, tt.wantBody) {
				// 	t.Errorf("complaintHandler.FindBetweenTimestamp() = %v, want %v", complaintList, tt.wantBody)
				// }

				/** 文字列同士でのアサーション */
				jsonByte, _ := json.Marshal(tt.wantBody)
				jsonStr := bytes.NewBuffer(jsonByte).String()
				fmt.Printf("w.Body.String(): %v\n", w.Body.String())
				fmt.Printf("jsonStr: %v\n", jsonStr)
				if !reflect.DeepEqual(w.Body.String(), jsonStr) {
					t.Errorf("complaintHandler.FindBetweenTimestamp() = %v, want %v", w.Body.String(), jsonStr)
				}

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
			wantErr:    errorJson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c := testutils.SetupGinContext()
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.DeleteByComplaintId(c)
			// ステータスコードのアサーション
			if !reflect.DeepEqual(c, tt.wantStatus) {
				t.Errorf("complaintHandler.DeleteByComplaintId() = %v, want %v", w.Code, tt.wantStatus)
			}

			// 異常系のレスポンスJSONのアサーション
			if tt.wantErr != nil {
				/** 【パターン1】jsonを構造体にマッピングして構造体同士でアサーション*/
				// var errorStruct *errors.ErrorStruct
				// fmt.Printf("w.Body.Bytes(): %s", w.Body.String())
				// json.Unmarshal(w.Body.Bytes(), &errorStruct)
				// if !reflect.DeepEqual(errorStruct, tt.wantErr) {
				// 	t.Errorf("complaintHandler.DeleteByComplaintId() = %v, want %v", errorStruct, tt.wantErr)
				// }

				/** 【パターン2】jsonを文字列に、構造体も文字列にマッピングして。文字列同士でアサーション*/
				fmt.Printf("w.Body.Sring(): %s\n", w.Body.String())
				jsonByte, _ := json.Marshal(tt.wantErr)
				jsonStr := bytes.NewBuffer(jsonByte).String()
				fmt.Printf("jsonStr: %s\n", jsonStr)
				if !reflect.DeepEqual(w.Body.String(), jsonStr) {
					t.Errorf("complaintHandler.DeleteByComplaintId() = %v, want %v", w.Body.String(), jsonStr)
				}
			}
		})
	}
}
