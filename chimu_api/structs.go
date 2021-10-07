package chimu_api

import (
	"encoding/json"
	"fmt"
)

type ChimuCommonResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type BeatmapInfo struct {
	BeatmapId        int64   `json:"BeatmapId"`
	ParentSetId      int64   `json:"ParentSetId"`
	DiffName         string  `json:"DiffName"`
	FileMD5          string  `json:"FileMD5"`
	Mode             int     `json:"Mode"`
	BPM              float64 `json:"BPM"`
	AR               float64 `json:"AR"`
	OD               float64 `json:"OD"`
	CS               float64 `json:"CS"`
	HP               float64 `json:"HP"`
	TotalLength      int     `json:"TotalLength"`
	HitLength        int     `json:"HitLength"`
	PlayCount        int64   `json:"Playcount"`
	PassCount        int64   `json:"Passcount"`
	DifficultyRating float64 `json:"DifficultyRating"`
	OsuFile          string  `json:"OsuFile"`
	DownloadPath     string  `json:"DownloadPath"`
}

type BeatmapSetInfo struct {
	SetId            int64         `json:"SetId"`
	ChildrenBeatmaps []BeatmapInfo `json:"ChildrenBeatmaps"`
	RankedStatus     int           `json:"RankedStatus"`
	ApprovedDate     string        `json:"ApprovedDate"`
	LastUpdate       string        `json:"LastUpdate"`
	LastChecked      string        `json:"LastChecked"`
	Artist           string        `json:"Artist"`
	Title            string        `json:"Title"`
	Creator          string        `json:"Creator"`
	Source           string        `json:"Source"`
	Tags             string        `json:"Tags"`
	HasVideo         bool          `json:"HasVideo"`
	Genre            int           `json:"Genre"`
	Language         int           `json:"Language"`
	Favourites       int64         `json:"Favourites"`
	Disabled         bool          `json:"Disabled"`
}

func (bsi *BeatmapSetInfo) FormatBeatmapSetName() (name string) {
	if len(bsi.Source) > 0 {
		name = fmt.Sprintf("%s (%s) - %s", bsi.Source, bsi.Artist, bsi.Title)
	} else {
		name = fmt.Sprintf("%s - %s", bsi.Artist, bsi.Title)
	}
	return
}
