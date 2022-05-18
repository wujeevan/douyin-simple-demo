package repository

import "testing"

func TestDoFavoriteVideo(t *testing.T) {
	type args struct {
		userId  int64
		videoId int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				userId:  8,
				videoId: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NotFavoriteVideo(tt.args.userId, tt.args.videoId); (err != nil) != tt.wantErr {
				t.Errorf("DoFavoriteVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
