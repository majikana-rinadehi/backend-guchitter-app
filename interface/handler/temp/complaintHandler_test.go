package handler

import (
	"reflect"
	"testing"

	"github.com/backend-guchitter-app/usecase"
	"github.com/gin-gonic/gin"
)

// TODO: setup mock of dependent layer

func TestNewComplaintHandler(t *testing.T) {
	// TODO: setup mock of dependent layer
	type args struct {
		cu usecase.ComplaintUseCase
	}
	tests := []struct {
		name string
		args args
		want ComplaintHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComplaintHandler(tt.args.cu); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComplaintHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complaintHandler_Index(t *testing.T) {
	// TODO: setup mock of dependent layer
	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.Index(tt.args.c)
		})
	}
}

func Test_complaintHandler_Search(t *testing.T) {
	// TODO: setup mock of dependent layer
	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.Search(tt.args.c)
		})
	}
}

func Test_complaintHandler_Create(t *testing.T) {
	// TODO: setup mock of dependent layer
	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.Create(tt.args.c)
		})
	}
}

func Test_complaintHandler_FindBetweenTimestamp(t *testing.T) {
	// TODO: setup mock of dependent layer
	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.FindBetweenTimestamp(tt.args.c)
		})
	}
}

func Test_complaintHandler_DeleteByComplaintId(t *testing.T) {
	// TODO: setup mock of dependent layer
	type fields struct {
		complaintUseCase usecase.ComplaintUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := complaintHandler{
				complaintUseCase: tt.fields.complaintUseCase,
			}
			ch.DeleteByComplaintId(tt.args.c)
		})
	}
}
