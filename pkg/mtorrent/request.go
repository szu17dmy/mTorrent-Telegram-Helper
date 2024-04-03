package mtorrent

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	EnvKeyMTorrentEndpoint         = "MTORRENT_ENDPOINT"
	EnvKeyMTorrentSecretKey        = "MTORRENT_SECRETKEY"
	RequestHeaderSecretKey         = "x-api-key"
	UrlMTorrentSecretKeyManagement = "https://kp.m-team.cc/usercp?tab=laboratory"
)

var (
	client       = &http.Client{}
	endpoint     = "https://test2.m-team.cc"
	secretKey    = ""
	HttpStatusOk = fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK))
)

type CommonResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func postForm(api string, data *map[string]string) (string, []byte, error) {
	formData := url.Values{}
	for key, value := range *data {
		formData.Set(key, value)
	}
	return post(api, formData.Encode(), "application/x-www-form-urlencoded")
}

func postJson(api string, form any) (string, []byte, error) {
	body, err := json.Marshal(form)
	if err != nil {
		return "", nil, err
	}
	return post(api, string(body), "application/json")
}

func post(api, body, contentType string) (string, []byte, error) {
	req, err := http.NewRequest("POST", endpoint+api, strings.NewReader(body))
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("Content-Type", contentType)
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
