package mtorrent

import (
	"log"
	"net/url"
	"os"
)

const (
	EnvKeyMTorrentEndpoint = "MTORRENT_ENDPOINT"
)

var (
	endpoint = "https://test2.m-team.cc"
)

func init() {
	initEndpoint()
}

func initEndpoint() {
	edp := os.Getenv(EnvKeyMTorrentEndpoint)
	if edp == "" {
		return
	}
	if _, err := url.Parse(edp); err != nil {
		log.Printf("failed to parse mTorrent endpoint: %s, use default: %s.", edp, endpoint)
		return
	}
	log.Printf("use mTorrent endpoint from env: %s", edp)
	endpoint = edp
}
