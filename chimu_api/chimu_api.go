package chimu_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	chimuURL = "https://api.chimu.moe"
)

func GetBeatmapV1(mapId int64) (info *BeatmapInfo, err error) {
	url := fmt.Sprintf("%s/v1/map/%d", chimuURL, mapId)

	body, _, statusCode, err := chimuAPIGet(url)
	if err != nil {
		return
	}
	resp := &ChimuCommonResponse{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	switch statusCode {
	case http.StatusNotFound:
		err = fmt.Errorf(resp.Message)
		return
	case http.StatusOK:
		err = json.Unmarshal(resp.Data, &info)
		return
	default:
		err = fmt.Errorf("unexpected http status code %d", statusCode)
		return
	}
}

func GetBeatmapSetV1(setId int64) (info *BeatmapSetInfo, err error) {
	url := fmt.Sprintf("%s/v1/set/%d", chimuURL, setId)

	body, _, statusCode, err := chimuAPIGet(url)
	if err != nil {
		return
	}
	resp := &ChimuCommonResponse{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	switch statusCode {
	case http.StatusNotFound:
		err = fmt.Errorf(resp.Message)
		return
	case http.StatusOK:
		err = json.Unmarshal(resp.Data, &info)
		return
	default:
		err = fmt.Errorf("unexpected http status code %d", statusCode)
		return
	}
}

func GetBeatmapSetDownloadURL(setId int64, noVideo bool) (dlUrl string, err error) {
	url := fmt.Sprintf("%s/v1/download/%d?n=", chimuURL, setId)
	if noVideo {
		url += "1"
	} else {
		url += "0"
	}

	body, header, statusCode, err := chimuAPIGet(url)
	if err != nil {
		return
	}

	switch statusCode {
	case http.StatusFound, http.StatusTemporaryRedirect:
		if len(header["Location"]) == 0 {
			err = fmt.Errorf("header Location not found")
			return
		}
		dlUrl = header["Location"][0]
		return
	case http.StatusForbidden, http.StatusUnauthorized, http.StatusNotFound:
		resp := &ChimuCommonResponse{}

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return
		}
		err = fmt.Errorf(resp.Message)
		return
	default:
		err = fmt.Errorf("unexpected http status code %d", statusCode)
		return
	}
}

func chimuAPIGet(url string) (body []byte, header http.Header, statusCode int, err error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	header = resp.Header

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
