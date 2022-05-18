package repository

import (
	"reflect"
	"testing"
)

func TestQueryFollowList(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*User
		wantErr bool
	}{
		{
			args: args{
				userId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryFollowList(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFollowList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryFollowList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryFollowerList(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*User
		wantErr bool
	}{
		{
			args: args{
				userId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryFollowerList(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFollowerList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryFollowerList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoFollowUser(t *testing.T) {
	type args struct {
		userId         int64
		followedUserId int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				userId:         2,
				followedUserId: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DoFollowUser(tt.args.userId, tt.args.followedUserId); (err != nil) != tt.wantErr {
				t.Errorf("DoFollowUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
