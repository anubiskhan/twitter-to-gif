package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"twitter-to-gif/services"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Please provide a URL")
		fmt.Println("Usage: vid-to-gif <url>")
		os.Exit(1)
	}

	inputURL := flag.Arg(0)
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		os.Exit(1)
	}

	// Get the appropriate service handler
	service := services.GetService(parsedURL.Host)
	if service == nil {
		fmt.Printf("Unsupported service: %s\n", parsedURL.Host)
		os.Exit(1)
	}

	// Create downloads directory if it doesn't exist
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	downloadDir := filepath.Join(homeDir, "Downloads")
	os.MkdirAll(downloadDir, 0755)

	// Process the URL
	outputPath, err := service.ProcessURL(inputURL, downloadDir)
	if err != nil {
		fmt.Printf("Error processing URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Media saved to: %s\n", outputPath)
}
