package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	mt "github.com/szu17dmy/mtorrent-telegram-helper/pkg/mtorrent"
)

func TestSendTorrentMessage(t *testing.T) {
	type args struct {
		torrent *mt.Torrent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sample",
			args: args{
				torrent: &mt.Torrent{
					Id:       "769649",
					Name:     "Post-graduate Entrance Examination",
					Abstract: "2023考研资料合集",
					Size:     "1393812756406",
					Status: mt.TorrentStatus{
						PinExpirationDate: "2024-04-08 00:00:00",
					},
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
		torrent *mt.Torrent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sample",
			args: args{
				torrent: &mt.Torrent{
					Id:       "769649",
					Name:     "Post-graduate Entrance Examination",
					Abstract: "2023考研资料合集",
					Size:     "1393812756406",
					Status: mt.TorrentStatus{
						PinExpirationDate: "2024-04-08 00:00:00",
					},
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
