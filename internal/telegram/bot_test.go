package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/ds"
	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/model"
)

func TestSendTorrentMessage(t *testing.T) {
	type args struct {
		torrent *model.Torrent
	}
	date, _ := ds.DateParse("2024-04-08 00:00:00")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sample",
			args: args{
				torrent: &model.Torrent{
					RemoteId:          "769649",
					Name:              "Post-graduate Entrance Examination",
					Abstract:          "2023考研资料合集",
					Size:              "1393812756406",
					PinExpirationDate: date,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendTorrentMessage(tt.args.torrent)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendTorrentMessage() error = %v, wantErr %v", err, tt.wantErr)
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

func TestSendNSFWTorrentMessage(t *testing.T) {
	type args struct {
		torrent *model.Torrent
	}
	date, _ := ds.DateParse("2024-04-08 00:00:00")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sample",
			args: args{
				torrent: &model.Torrent{
					RemoteId:          "769649",
					Name:              "Post-graduate Entrance Examination",
					Abstract:          "2023考研资料合集",
					Size:              "1393812756406",
					PinExpirationDate: date,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendNSFWTorrentMessage(tt.args.torrent)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendNSFWTorrentMessage() error = %v, wantErr %v", err, tt.wantErr)
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
