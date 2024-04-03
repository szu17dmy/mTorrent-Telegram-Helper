package mtorrent

import (
	"log"
	"net/url"
	"os"
)

const (
	EnvKeyMTorrentEndpoint         = "MTORRENT_ENDPOINT"
	EnvKeyMTorrentSecretKey        = "MTORRENT_SECRETKEY"
	UrlMTorrentSecretKeyManagement = "https://kp.m-team.cc/usercp?tab=laboratory"
)

var (
	endpoint  = "https://test2.m-team.cc"
	secretKey = ""
)

func post(api string, form any) {

}

func init() {
	initEndpoint()
	initSecretKey()
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

func initSecretKey() {
	sk := os.Getenv(EnvKeyMTorrentSecretKey)
	if sk == "" {
		log.Fatalf("failed to load secret key, please generate one from: " + UrlMTorrentSecretKeyManagement)
	}
	secretKey = sk
}
