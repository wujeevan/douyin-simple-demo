package service

import (
	"reflect"
	"testing"
)

func TestSignUp(t *testing.T) {
	SignUp("admin", "pass")
}

func TestSignIn(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				username: "xxxyyy",
				password: "hellowo",
			},
			want: "username not exists",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignIn(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
