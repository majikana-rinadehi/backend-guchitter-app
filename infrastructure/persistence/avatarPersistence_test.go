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
	wantAvatar = &model.Avatar{
		AvatarId:   1,
		AvatarName: "あああ",
		AvatarText: "あああ",
		ImageUrl:   "https://example.com",
		Color:      "#ffffff",
	}
	wantAvatarList = []*model.Avatar{
		{
			AvatarId:   1,
			AvatarName: "あああ",
			AvatarText: "あああ",
			ImageUrl:   "https://example.com",
			Color:      "#ffffff",
		},
		{
			AvatarId:   2,
			AvatarName: "いいい",
			AvatarText: "いいい",
			ImageUrl:   "https://example.com",
			Color:      "#ffffff",
		},
		{
			AvatarId:   3,
			AvatarName: "ううう",
			AvatarText: "ううう",
			ImageUrl:   "https://example.com",
			Color:      "#ffffff",
		},
	}
	insertAvatar = &model.Avatar{
		AvatarId:   4,
		AvatarName: "えええ",
		AvatarText: "えええ",
		ImageUrl:   "https://example.com",
		Color:      "#ffffff",
	}
)

func Test_avatarPersistence_FindAll(t *testing.T) {
	// setup env
	testUtils.SetTestEnv(t)

	// setup db
	testUtils.SetupFixtures()

	type fields struct {
		Conn *gorm.DB
	}
	tests := []struct {
		name           string
		fields         fields
		wantAvatarList []*model.Avatar
		wantErr        error
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			wantAvatarList: wantAvatarList,
			wantErr:        nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &avatarPersistence{
				Conn: tt.fields.Conn,
			}
			gotAvatarList, err := cp.FindAll()
			if err != nil && err != tt.wantErr {
				t.Errorf("avatarPersistence.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAvatarList, tt.wantAvatarList) {
				t.Errorf("avatarPersistence.FindAll() = %v, want %v", gotAvatarList, tt.wantAvatarList)
			}
		})
	}
}

func Test_avatarPersistence_FindByAvatarId(t *testing.T) {
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
		name       string
		fields     fields
		args       args
		wantAvatar *model.Avatar
		wantErr    error
	}{
		// TODO: Add test cases.
		{
			name: "Test_avatarPersistence_FindByAvatarId_Normal",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				id: 1,
			},
			wantAvatar: wantAvatar,
			wantErr:    nil,
		},
		{
			name: "Test_avatarPersistence_FindByAvatarId_NotFound",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				id: 999,
			},
			wantAvatar: nil,
			wantErr:    gorm.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &avatarPersistence{
				Conn: tt.fields.Conn,
			}
			gotAvatar, err := cp.FindByAvatarId(tt.args.id)
			if err != nil && err != tt.wantErr {
				t.Errorf("avatarPersistence.FindByAvatarId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAvatar, tt.wantAvatar) {
				t.Errorf("avatarPersistence.FindByAvatarId() = %v, want %v", gotAvatar, tt.wantAvatar)
			}
		})
	}
}

func Test_avatarPersistence_Create(t *testing.T) {
	// setup env
	testUtils.SetTestEnv(t)

	// setup db
	testUtils.SetupFixtures()

	type fields struct {
		Conn *gorm.DB
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
			name: "Test_avatarPersistence_Create_Normal",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				avatar: *insertAvatar,
			},
			want:    insertAvatar,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &avatarPersistence{
				Conn: tt.fields.Conn,
			}
			got, err := cp.Create(tt.args.avatar)
			if err != nil && err != tt.wantErr {
				t.Errorf("avatarPersistence.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avatarPersistence.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_avatarPersistence_FindBetweenTimestamp(t *testing.T) {
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
		name           string
		fields         fields
		args           args
		wantAvatarList []*model.Avatar
		wantErr        error
	}{
		// TODO: Add test cases.
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
			name: "Test_avatarPersistence_FindBetweenTimestamp_1-1",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				from: "2023-05-06 13:59:00",
			},
			wantAvatarList: []*model.Avatar{
				{AvatarId: 1, AvatarText: "あああ", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 2, AvatarText: "いいい", AvatarName: "いいい", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 3, AvatarText: "ううう", AvatarName: "ううう", ImageUrl: "https://example.com", Color: "#ffffff"},
			},
			wantErr: nil,
		},
		{
			// 1-2. from = last_update
			name: "Test_avatarPersistence_FindBetweenTimestamp_1-2",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				// from: "2023-05-06 14:00:00",
				from: "2023-05-06 23:00:00",
			},
			wantAvatarList: []*model.Avatar{
				{AvatarId: 1, AvatarText: "あああ", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 2, AvatarText: "いいい", AvatarName: "いいい", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 3, AvatarText: "ううう", AvatarName: "ううう", ImageUrl: "https://example.com", Color: "#ffffff"},
			},
			wantErr: nil,
		},
		{
			// 1-3. from > last_update
			name: "Test_avatarPersistence_FindBetweenTimestamp_1-3",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				// from: "2023-05-06 14:00:01",
				from: "2023-05-06 23:00:01",
			},
			wantAvatarList: []*model.Avatar{
				// {AvatarId: 1, AvatarText: "あああ", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff" },
				{AvatarId: 2, AvatarText: "いいい", AvatarName: "いいい", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 3, AvatarText: "ううう", AvatarName: "ううう", ImageUrl: "https://example.com", Color: "#ffffff"},
			},
			wantErr: nil,
		},
		{
			// 2-1. last_update < to
			name: "Test_avatarPersistence_FindBetweenTimestamp_2-1",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				to: "2023-05-07 1:00:01",
			},
			wantAvatarList: []*model.Avatar{
				{AvatarId: 1, AvatarText: "あああ", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 2, AvatarText: "いいい", AvatarName: "いいい", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 3, AvatarText: "ううう", AvatarName: "ううう", ImageUrl: "https://example.com", Color: "#ffffff"},
			},
			wantErr: nil,
		},
		{
			// 2-2. last_update = to
			name: "Test_avatarPersistence_FindBetweenTimestamp_2-2",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				to: "2023-05-07 1:00:00",
			},
			wantAvatarList: []*model.Avatar{
				{AvatarId: 1, AvatarText: "あああ", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 2, AvatarText: "いいい", AvatarName: "いいい", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 3, AvatarText: "ううう", AvatarName: "ううう", ImageUrl: "https://example.com", Color: "#ffffff"},
			},
			wantErr: nil,
		},
		{
			// 2-3. last_update > to
			name: "Test_avatarPersistence_FindBetweenTimestamp_2-3",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				to: "2023-05-07 0:59:59",
			},
			wantAvatarList: []*model.Avatar{
				{AvatarId: 1, AvatarText: "あああ", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff"},
				{AvatarId: 2, AvatarText: "いいい", AvatarName: "いいい", ImageUrl: "https://example.com", Color: "#ffffff"},
				// {AvatarId: 3, AvatarText: "ううう", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff" },
			},
			wantErr: nil,
		},
		{
			// 3-1. from < last_update < to
			name: "Test_avatarPersistence_FindBetweenTimestamp_3-1",
			fields: fields{
				Conn: config.ConnectTest(),
			},
			args: args{
				from: "2023-05-06 23:00:01",
				to:   "2023-05-07 0:59:59",
			},
			wantAvatarList: []*model.Avatar{
				// {AvatarId: 1, AvatarText: "あああ", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff" },
				{AvatarId: 2, AvatarText: "いいい", AvatarName: "いいい", ImageUrl: "https://example.com", Color: "#ffffff"},
				// {AvatarId: 3, AvatarText: "ううう", AvatarName: "あああ", ImageUrl: "https://example.com", Color: "#ffffff" },
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &avatarPersistence{
				Conn: tt.fields.Conn,
			}
			gotAvatarList, err := cp.FindBetweenTimestamp(tt.args.from, tt.args.to)
			if err != nil && err != tt.wantErr {
				t.Errorf("avatarPersistence.FindBetweenTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAvatarList, tt.wantAvatarList) {
				t.Errorf("avatarPersistence.FindBetweenTimestamp() = %v, want %v", gotAvatarList, tt.wantAvatarList)
			}
		})
	}
}

func Test_avatarPersistence_DeleteByAvatarId(t *testing.T) {

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
			name: "Test_avatarPersistence_DeleteByAvatarId_Normal",
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
			cp := &avatarPersistence{
				Conn: tt.fields.Conn,
			}
			if err := cp.DeleteByAvatarId(tt.args.id); err != nil && err != tt.wantErr {
				t.Errorf("avatarPersistence.DeleteByAvatarId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
