package mtorrent

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestTorrentSearch(t *testing.T) {
	type args struct {
		req *TorrentSearchRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *TorrentSearchResponse
		wantErr bool
	}{
		{
			name: "Search",
			args: args{
				req: &TorrentSearchRequest{
					PageNumber: 1,
					PageSize:   10,
					Mode:       "Normal",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TorrentSearch(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TorrentSearch() error = %v, wantErr %v", err, tt.wantErr)
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
