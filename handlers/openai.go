package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// GenerateTarotCardDesign generates a tarot card design using DALL·E
func GenerateTarotCardDesign(card, theme, color1 string, color2 string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		// If no API key is provided, return a default response
		return "No API key provided", nil
	}

	url := os.Getenv("OPENAI_BASE_URL") // Use the DALL·E image generation endpoint

	// Build the request body
	requestBody, err := buildPostRequestBody(card, theme, color1, color2)
	if err != nil {
		return "", fmt.Errorf("error building request body: %w", err)
	}

	// Create and send the HTTP request
	req, err := createRequest(url, apiKey, requestBody)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	responseBody, err := sendRequest(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to API endpoint: %w", err)
	}

	// Parse and return the image URL
	return parseImageResponse(responseBody)
}

func buildPostRequestBody(card, theme, color1 string, color2 string) ([]byte, error) {
	if card == "" || theme == "" || color1 == "" || color2 == "" {
		return nil, fmt.Errorf("card, theme, and color are required")
	}

	// Construct the prompt for DALL·E image generation
	prompt := fmt.Sprintf(
		"Design a detailed illustration for the front side of the tarot card '%s', in a '%s' style. The artwork should fill the entire canvas with no borders or empty spaces, emphasizing hues of '%s' and secondary hues of '%s'. Incorporate traditional tarot symbolism and ensure the design is vertically oriented. Do not add numbers or letters whatsoever in the design.",
		card, theme, color1, color2,
	)

	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt":  prompt,
		"model":   "dall-e-3",
		"n":       1,
		"size":    "1024x1792",
		"quality": "hd",
	})
	if err != nil {
		return nil, err
	}

	return requestBody, nil
}

func createRequest(url, apiKey string, requestBody []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func sendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Print or log the full response body for debugging
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Response Body: ", string(body)) // For debugging
	return body, nil
}

// parseImageResponse parses the response from DALL·E and returns the image URL
func parseImageResponse(responseBody []byte) (string, error) {
	var response map[string]interface{}
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		return "", err
	}

	// Extract the data field
	data, ok := response["data"].([]interface{})
	if !ok || len(data) == 0 {
		return "", fmt.Errorf("no images returned in the response")
	}

	// Extract the first image URL
	firstImage, ok := data[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid image data")
	}

	imageURL, ok := firstImage["url"].(string)
	if !ok {
		return "", fmt.Errorf("no URL found in the response")
	}

	return imageURL, nil
}
