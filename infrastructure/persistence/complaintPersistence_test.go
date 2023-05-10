package persistence

import (
	"reflect"
	"testing"

	"github.com/backend-guchitter-app/config"
	"github.com/backend-guchitter-app/domain/model"
	testUtils "github.com/backend-guchitter-app/util/testUtils"
	"gorm.io/gorm"
)

var (
	wantComplaint = &model.Complaint{
		ComplaintId:   1,
		ComplaintText: "あああ",
		AvatarId:      1,
	}
	wantComplaintList = []*model.Complaint{
		{
			ComplaintId:   1,
			ComplaintText: "あああ",
			AvatarId:      1,
		},
		{
			ComplaintId:   2,
			ComplaintText: "いいい",
			AvatarId:      1,
		},
		{
			ComplaintId:   3,
			ComplaintText: "ううう",
			AvatarId:      1,
		},
	}
	insertComplaint = &model.Complaint{
		ComplaintId:   4,
		ComplaintText: "えええ",
		AvatarId:      1,
	}
)

func Test_complaintPersistence_FindAll(t *testing.T) {

	// setup env
	testUtils.SetTestEnv(t)

	// setup db
	testUtils.SetupFixtures()

	type fields struct {
		Conn *gorm.DB
	}
	tests := []struct {
		name              string
		fields            fields
		wantComplaintList []*model.Complaint
		wantErr           error
	}{
		// TODO: Add test cases.
		{
			name: "Test_complaintPersistence_FindAll_Normal",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			wantComplaintList: wantComplaintList,
			wantErr:           nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &complaintPersistence{
				Conn: tt.fields.Conn,
			}
			gotComplaintList, err := cp.FindAll()
			if err != nil && err != tt.wantErr {
				t.Errorf("complaintPersistence.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotComplaintList, tt.wantComplaintList) {
				t.Errorf("complaintPersistence.FindAll() = %v, want %v",
					testUtils.ToString(gotComplaintList), testUtils.ToString(tt.wantComplaintList))
			}
		})
	}
}

func Test_complaintPersistence_FindByAvatarId(t *testing.T) {
	// setup env
	testUtils.SetTestEnv(t)

	// setup db
	testUtils.SetupFixtures()

	type fields struct {
		Conn *gorm.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantComplaint *model.Complaint
		wantErr       error
	}{
		// TODO: Add test cases.
		{
			name: "Test_complaintPersistence_FindByAvatarId_Normal",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				id: 1,
			},
			wantComplaint: wantComplaint,
			wantErr:       nil,
		},
		{
			name: "Test_complaintPersistence_FindByAvatarId_NotFound",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				id: 999,
			},
			wantComplaint: nil,
			wantErr:       gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &complaintPersistence{
				Conn: tt.fields.Conn,
			}
			gotComplaint, err := cp.FindByAvatarId(tt.args.id)
			if err != nil && err != tt.wantErr {
				t.Errorf("complaintPersistence.FindByAvatarId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotComplaint, tt.wantComplaint) {
				t.Errorf("complaintPersistence.FindByAvatarId() = %v, want %v",
					testUtils.ToString(gotComplaint), testUtils.ToString(tt.wantComplaint))
			}
		})
	}
}

func Test_complaintPersistence_Create(t *testing.T) {
	// setup env
	testUtils.SetTestEnv(t)

	// setup db
	testUtils.SetupFixtures()

	type fields struct {
		Conn *gorm.DB
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
			name: "Test_complaintPersistence_Create_Normal",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				complaint: *insertComplaint,
			},
			want:    insertComplaint,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &complaintPersistence{
				Conn: tt.fields.Conn,
			}
			got, err := cp.Create(tt.args.complaint)
			if err != nil && err != tt.wantErr {
				t.Errorf("complaintPersistence.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("complaintPersistence.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complaintPersistence_FindBetweenTimestamp(t *testing.T) {
	// setup env
	testUtils.SetTestEnv(t)

	// setup db
	testUtils.SetupFixtures()

	type fields struct {
		Conn *gorm.DB
	}
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantComplaintList []*model.Complaint
		wantErr           error
	}{
		// TODO: Add test cases.
		// 1-1. from < last_update
		// 1-2. from = last_update
		// 1-3. from > last_update
		// 2-1. last_update < to
		// 2-2. last_update = to
		// 2-3. last_update > to
		// 3-1. from < last_update < to
		{
			// 1-1. from < last_update
			name: "Test_complaintPersistence_FindBetweenTimestamp_1-1",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				from: "2023-05-06 13:59:00",
			},
			wantComplaintList: []*model.Complaint{
				{ComplaintId: 1, ComplaintText: "あああ", AvatarId: 1},
				{ComplaintId: 2, ComplaintText: "いいい", AvatarId: 1},
				{ComplaintId: 3, ComplaintText: "ううう", AvatarId: 1},
			},
			wantErr: nil,
		},
		{
			// 1-2. from = last_update
			name: "Test_complaintPersistence_FindBetweenTimestamp_1-2",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				// from: "2023-05-06 14:00:00",
				from: "2023-05-06 23:00:00",
			},
			wantComplaintList: []*model.Complaint{
				{ComplaintId: 1, ComplaintText: "あああ", AvatarId: 1},
				{ComplaintId: 2, ComplaintText: "いいい", AvatarId: 1},
				{ComplaintId: 3, ComplaintText: "ううう", AvatarId: 1},
			},
			wantErr: nil,
		},
		{
			// 1-3. from > last_update
			name: "Test_complaintPersistence_FindBetweenTimestamp_1-3",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				// from: "2023-05-06 14:00:01",
				from: "2023-05-06 23:00:01",
			},
			wantComplaintList: []*model.Complaint{
				// {ComplaintId: 1, ComplaintText: "あああ", AvatarId: 1},
				{ComplaintId: 2, ComplaintText: "いいい", AvatarId: 1},
				{ComplaintId: 3, ComplaintText: "ううう", AvatarId: 1},
			},
			wantErr: nil,
		},
		{
			// 2-1. last_update < to
			name: "Test_complaintPersistence_FindBetweenTimestamp_2-1",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				to: "2023-05-07 1:00:01",
			},
			wantComplaintList: []*model.Complaint{
				{ComplaintId: 1, ComplaintText: "あああ", AvatarId: 1},
				{ComplaintId: 2, ComplaintText: "いいい", AvatarId: 1},
				{ComplaintId: 3, ComplaintText: "ううう", AvatarId: 1},
			},
			wantErr: nil,
		},
		{
			// 2-2. last_update = to
			name: "Test_complaintPersistence_FindBetweenTimestamp_2-2",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				to: "2023-05-07 1:00:00",
			},
			wantComplaintList: []*model.Complaint{
				{ComplaintId: 1, ComplaintText: "あああ", AvatarId: 1},
				{ComplaintId: 2, ComplaintText: "いいい", AvatarId: 1},
				{ComplaintId: 3, ComplaintText: "ううう", AvatarId: 1},
			},
			wantErr: nil,
		},
		{
			// 2-3. last_update > to
			name: "Test_complaintPersistence_FindBetweenTimestamp_2-3",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				to: "2023-05-07 0:59:59",
			},
			wantComplaintList: []*model.Complaint{
				{ComplaintId: 1, ComplaintText: "あああ", AvatarId: 1},
				{ComplaintId: 2, ComplaintText: "いいい", AvatarId: 1},
				// {ComplaintId: 3, ComplaintText: "ううう", AvatarId: 1},
			},
			wantErr: nil,
		},
		{
			// 3-1. from < last_update < to
			name: "Test_complaintPersistence_FindBetweenTimestamp_3-1",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				from: "2023-05-06 23:00:01",
				to:   "2023-05-07 0:59:59",
			},
			wantComplaintList: []*model.Complaint{
				// {ComplaintId: 1, ComplaintText: "あああ", AvatarId: 1},
				{ComplaintId: 2, ComplaintText: "いいい", AvatarId: 1},
				// {ComplaintId: 3, ComplaintText: "ううう", AvatarId: 1},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &complaintPersistence{
				Conn: tt.fields.Conn,
			}
			gotComplaintList, err := cp.FindBetweenTimestamp(tt.args.from, tt.args.to)
			if err != nil && err != tt.wantErr {
				t.Errorf("complaintPersistence.FindBetweenTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotComplaintList, tt.wantComplaintList) {
				t.Errorf("complaintPersistence.FindBetweenTimestamp() = %v, want %v",
					testUtils.ToString(gotComplaintList), testUtils.ToString(tt.wantComplaintList))
			}
		})
	}
}

func Test_complaintPersistence_DeleteByComplaintId(t *testing.T) {
	// setup env
	testUtils.SetTestEnv(t)

	// setup db
	testUtils.SetupFixtures()

	type fields struct {
		Conn *gorm.DB
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
			name: "Test_complaintPersistence_DeleteByComplaintId_Normal",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				id: 1,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &complaintPersistence{
				Conn: tt.fields.Conn,
			}
			if err := cp.DeleteByComplaintId(tt.args.id); err != nil && err != tt.wantErr {
				t.Errorf("complaintPersistence.DeleteByComplaintId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
