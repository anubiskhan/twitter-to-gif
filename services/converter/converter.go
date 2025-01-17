package converter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Converter struct {
	containerImage string
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

func (c *Converter) ConvertToGif(inputURL string, outputDir string) (string, error) {
	var outb bytes.Buffer
	cmd := exec.Command("docker", "run", "--rm",
		c.containerImage,
		inputURL,
	)

	cmd.Stdout = &outb
	cmd.Stderr = os.Stderr

	fmt.Printf("Running converter for URL: %s\n", inputURL)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("conversion failed: %v", err)
	}

	// Parse the JSON response
	var response ConversionResponse
	if err := json.NewDecoder(&outb).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to parse converter response: %v", err)
	}

	if response.Error != "" {
		return "", fmt.Errorf("converter error: %s", response.Error)
	}

	fmt.Printf("Received filename: %s\n", response.Filename)

	// Decode base64 data and write to file
	data, err := base64.StdEncoding.DecodeString(response.Data)
	if err != nil {
		return "", fmt.Errorf("failed to decode gif data: %v", err)
	}

	// Write to the mounted workdir
	outputPath := filepath.Join("/workdir", response.Filename)
	fmt.Printf("Writing to path: %s\n", outputPath)

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return "", fmt.Errorf("failed to write gif file: %v", err)
	}

	hostPath := filepath.Join(outputDir, response.Filename)
	fmt.Printf("Returning host path: %s\n", hostPath)

	// Return the path relative to the host's Downloads directory
	return hostPath, nil
}
