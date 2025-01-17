package services

import (
	"strings"
)

type DownloadMode int

const (
	ModeGIF DownloadMode = iota
	ModeVideo
)

type MediaService interface {
	ProcessURL(url string, outputDir string, mode DownloadMode) (string, error)
}

func GetService(hostname string) MediaService {
	switch {
	case strings.Contains(hostname, "twitter.com"), strings.Contains(hostname, "x.com"):
		return NewTwitterService()
	case strings.Contains(hostname, "instagram.com"):
		return NewInstagramService()
	default:
		return nil
	}
}
