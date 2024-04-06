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
