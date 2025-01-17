package services

import (
	"twitter-to-gif/services/converter"
)

type InstagramService struct {
	converter *converter.Converter
}

func NewInstagramService() *InstagramService {
	return &InstagramService{
		converter: converter.New(),
	}
}

func (s *InstagramService) ProcessURL(url string, outputDir string) (string, error) {
	return s.converter.ConvertToGif(url, outputDir)
}
