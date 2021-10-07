package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/fubuki-is-cat/chimu_client/chimu_api"
	log "github.com/sirupsen/logrus"
)

var (
	args struct {
		NoVideo bool
		URL     string
	}
)

const _VERSION = "1.0.0"

func init() {
	flag.BoolVar(&args.NoVideo, "no-video", false, "download without video")
	flag.StringVar(&args.URL, "url", "", "beatmap url, e.g: https://osu.ppy.sh/beatmaps/1540798")
	flag.Parse()
}

func main() {
	log.WithField("version", _VERSION).Infoln("Chimu client started")
	if len(args.URL) == 0 {
		log.Fatalln("Beatmap URL is empty")
	}

	logger := log.WithField("url", args.URL)

	splURL := strings.Split(args.URL, "/")
	mapId, err := strconv.ParseInt(splURL[len(splURL)-1], 10, 64)
	if err != nil {
		logger.Fatalln(err)
	}
	logger = log.WithField("beatmap_id", mapId)

	logger.Infoln("get beatmap info")
	info, err := chimu_api.GetBeatmapV1(mapId)
	if err != nil {
		logger.Fatalln(err)
	}
	logger.WithFields(log.Fields{
		"beatmap_id":    info.BeatmapId,
		"parent_set_id": info.ParentSetId,
		"diff_name":     info.DiffName,
		"filename":      info.OsuFile,
	}).Infoln("beatmap info")

	logger = log.WithField("beatmap_set_id", info.ParentSetId)
	logger.Infoln("get beatmap set info")
	setInfo, err := chimu_api.GetBeatmapSetV1(info.ParentSetId)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.WithFields(log.Fields{
		"beatmap_set_id": setInfo.SetId,
		"ranked_status":  setInfo.RankedStatus,
		"creator":        setInfo.Creator,
	}).Infoln(setInfo.FormatBeatmapSetName())

	logger.Infoln("get download link")
	dlUrl, err := chimu_api.GetBeatmapSetDownloadURL(info.ParentSetId, args.NoVideo)
	if err != nil {
		logger.Fatalln(err)
	}
	logger.WithField("dl_url", dlUrl).Infoln("downloading beatmap set")
	fn, err := downloadBeatmap(dlUrl)
	if err != nil {
		log.Fatalln(err)
	}
	logger.WithField("filename", fn).Infoln("beatmap set downloaded")
}
