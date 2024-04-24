package mtorrent

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	mt "github.com/szu17dmy/mtorrent-telegram-helper/pkg/mtorrent"
)

func TestGetFreeAndLargeNormalTorrents(t *testing.T) {
	tests := []struct {
		name    string
		want    []*mt.Torrent
		wantErr bool
	}{
		{
			name: "Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFreeAndLargeNormalTorrents()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFreeAndLargeNormalTorrents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			bs, err := json.MarshalIndent(got, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(bs))
		})
	}
}

func Test_dateAfterNow(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				date: "null",
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				date: "2023-01-01 00:00:00",
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				date: "2099-01-01 00:00:00",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dateAfterNow(tt.args.date); got != tt.want {
				t.Errorf("dateAfterNow() = %v, want %v", got, tt.want)
			}
		})
	}
}
