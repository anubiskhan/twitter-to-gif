package services

import (
	"twitter-to-gif/services/converter"
)

type TwitterService struct {
	converter *converter.Converter
}

func NewTwitterService() *TwitterService {
	return &TwitterService{
		converter: converter.New(),
	}
}

func (s *TwitterService) ProcessURL(url string, outputDir string) (string, error) {
	return s.converter.ConvertToGif(url, outputDir)
}
