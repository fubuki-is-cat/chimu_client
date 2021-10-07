package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/dustin/go-humanize"
)

type dlCounter struct {
	ReceivedSize uint64
	TotalSize    uint64
}

func (wc *dlCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.ReceivedSize += uint64(n)
	fmt.Printf("\rDownloading... %s/%s (%d%%)", humanize.Bytes(wc.ReceivedSize), humanize.Bytes(wc.TotalSize), wc.ReceivedSize*100/wc.TotalSize)
	return n, nil
}

func downloadBeatmap(dlUrl string) (fn string, err error) {
	uParse, err := url.Parse(dlUrl)
	if err != nil {
		return
	}
	query := uParse.Query()
	fn = query.Get("filename")

	out, err := os.Create(fn + ".tmp")
	if err != nil {
		return
	}

	resp, err := http.Get(dlUrl)
	if err != nil {
		out.Close()
		return
	}
	defer resp.Body.Close()

	counter := &dlCounter{TotalSize: uint64(resp.ContentLength)}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return
	}
	fmt.Print("\r")
	out.Close()

	if err = os.Rename(fn+".tmp", fn); err != nil {
		return
	}
	return
}
