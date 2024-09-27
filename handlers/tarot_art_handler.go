package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// downloadImage downloads the tarot card image from the provided URL and saves it in the /public/images folder.
// It checks if the file already exists and appends a number to the file name if needed.
func DownloadImage(url, cardName string) (string, error) {
	// Define the base path where the image will be saved
	basePath := "./public/images"
	extension := ".png"
	fileName := cardName + extension
	filePath := filepath.Join(basePath, fileName)

	// Check if a file with the same name already exists, and if so, modify the file name
	filePath, err := ensureUniqueFileName(basePath, cardName, extension)
	if err != nil {
		return "", err
	}

	// Download the image from the URL
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Copy the downloaded content to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	return filePath, nil
}

// ensureUniqueFileName checks if a file with the same name already exists in the folder.
// If it exists, it appends a number (e.g., -2, -3, etc.) to the file name.
func ensureUniqueFileName(basePath, cardName, extension string) (string, error) {
	filePath := filepath.Join(basePath, cardName+extension)

	// If the file does not exist, return the original file path
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return filePath, nil
	}

	// If the file exists, append a number to the file name
	counter := 2
	for {
		modifiedFileName := fmt.Sprintf("%s-%d%s", cardName, counter, extension)
		modifiedFilePath := filepath.Join(basePath, modifiedFileName)

		// Check if the modified file name exists
		if _, err := os.Stat(modifiedFilePath); os.IsNotExist(err) {
			return modifiedFilePath, nil
		}

		// Increment the counter and try again
		counter++
	}
}
