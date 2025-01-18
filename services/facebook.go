package services

import (
	"os"
	"path/filepath"
	"twitter-to-gif/services/converter"
)

type FacebookService struct {
	converter *converter.Converter
}

func NewFacebookService() *FacebookService {
	return &FacebookService{
		converter: converter.New(),
	}
}

func (s *FacebookService) ProcessURL(url string, outputDir string, mode DownloadMode) (string, error) {
	// Convert mode to converter mode
	convMode := converter.ModeGIF
	if mode == ModeVideo {
		convMode = converter.ModeVideo
	}

	// Get the media data from converter
	filename, data, err := s.converter.Convert(url, convMode)
	if err != nil {
		return "", err
	}

	// Write to the mounted workdir
	outputPath := filepath.Join("/workdir", filename)
	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return "", err
	}

	// Return the path relative to the host's Downloads directory
	return filepath.Join(outputDir, filename), nil
}
