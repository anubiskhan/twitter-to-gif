package converter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Converter struct {
	containerImage string
}

type ConversionMode string

const (
	ModeGIF   ConversionMode = "gif"
	ModeVideo ConversionMode = "video"
)

type ConversionRequest struct {
	URL  string         `json:"url"`
	Mode ConversionMode `json:"mode"`
}

type ConversionResponse struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
	Error    string `json:"error,omitempty"`
}

func New() *Converter {
	return &Converter{
		containerImage: "twitter-to-gif-converter:latest",
	}
}

func (c *Converter) Convert(inputURL string, mode ConversionMode) (string, []byte, error) {
	request := ConversionRequest{
		URL:  inputURL,
		Mode: mode,
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return "", nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	var outb bytes.Buffer
	cmd := exec.Command("docker", "run", "--rm",
		"-i", // Enable stdin for JSON input
		c.containerImage,
	)

	cmd.Stdin = bytes.NewReader(requestJSON)
	cmd.Stdout = &outb
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", nil, fmt.Errorf("conversion failed: %v", err)
	}

	var response ConversionResponse
	if err := json.NewDecoder(&outb).Decode(&response); err != nil {
		return "", nil, fmt.Errorf("failed to parse converter response: %v", err)
	}

	if response.Error != "" {
		return "", nil, fmt.Errorf("converter error: %s", response.Error)
	}

	data, err := base64.StdEncoding.DecodeString(response.Data)
	if err != nil {
		return "", nil, fmt.Errorf("failed to decode media data: %v", err)
	}

	return response.Filename, data, nil
}
