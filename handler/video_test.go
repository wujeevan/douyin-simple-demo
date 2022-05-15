package handler

import (
	"os"
	"reflect"
	"testing"

	"github.com/wujeevan/douyinv0/repository"
)

func TestMain(m *testing.M) {
	if err := repository.Init(); err != nil {
		os.Exit(1)
	}
	m.Run()
}

func TestQueryFeedVideo(t *testing.T) {
	type args struct {
		latest_time string
		token       string
	}
	tests := []struct {
		name string
		args args
		want *FeedVideoResponse
	}{
		{
			args: args{
				latest_time: "1652315460000",
				token:       "abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryFeedVideo(tt.args.latest_time, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryFeedVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}
