package mtorrent

import (
	"encoding/json"
	"errors"
)

const (
	apiTorrentSearch = "/api/torrent/search"
)

type Torrent struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Abstract string        `json:"smallDescr"`
	Size     string        `json:"size"`
	Status   TorrentStatus `json:"status"`
}

type TorrentStatus struct {
	Discount               string `json:"discount"`
	DiscountExpirationDate string `json:"discountEndTime"`
	Pin                    int    `json:"toppingLevel"`
	PinExpirationDate      string `json:"toppingEndTime"`
}

type TorrentSearchRequest struct {
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	Keyword    string `json:"keyword"`
	Visible    int    `json:"visible"`
	Mode       string `json:"mode"`
}

type TorrentSearchResponse struct {
	CommonResponse
	Data TorrentSearchResponseData `json:"data"`
}

type TorrentSearchResponseData struct {
	Data []*Torrent
}

func TorrentSearch(req *TorrentSearchRequest) (*TorrentSearchResponse, error) {
	stat, resp, err := postJson(apiTorrentSearch, req)
	if err != nil {
		return nil, err
	}
	if stat != HttpStatusOk {
		return nil, errors.New(stat + string(resp))
	}
	ent := &TorrentSearchResponse{}
	err = json.Unmarshal(resp, ent)
	return ent, err
}
