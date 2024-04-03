package mtorrent

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	EnvKeyMTorrentEndpoint         = "MTORRENT_ENDPOINT"
	EnvKeyMTorrentSecretKey        = "MTORRENT_SECRETKEY"
	RequestHeaderSecretKey         = "x-api-key"
	UrlMTorrentSecretKeyManagement = "https://kp.m-team.cc/usercp?tab=laboratory"
)

var (
	client    = &http.Client{}
	endpoint  = "https://test2.m-team.cc"
	secretKey = ""
)

func post(api string, form any) (string, []byte, error) {
	body, err := json.Marshal(form)
	if err != nil {
		return "", nil, err
	}
	req, err := http.NewRequest("POST", endpoint+api, bytes.NewBuffer(body))
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(RequestHeaderSecretKey, secretKey)
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}
	return resp.Status, res, nil
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
