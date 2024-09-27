package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// GenerateTarotCardDesign generates a tarot card based on the given card, theme, and color using the OpenAI API.
// It returns the image URL of the generated card.
func GenerateTarotCardDesign(card, theme, color1 string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key not provided")
	}

	url := os.Getenv("OPENAI_BASE_URL")
	if url == "" {
		return "", fmt.Errorf("API base URL not provided")
	}

	// Build the POST request body
	requestBody, err := buildPostRequestBody(card, theme, color1)
	if err != nil {
		return "", fmt.Errorf("error building request body: %w", err)
	}

	// Create the request
	req, err := createRequest(url, apiKey, requestBody)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Send the request and handle the response
	responseBody, err := sendRequest(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to API endpoint: %w", err)
	}

	// Parse the image URL from the response
	imageURL, err := parseImageResponse(responseBody)
	if err != nil {
		return "", fmt.Errorf("error parsing image response: %w", err)
	}

	return imageURL, nil
}

// buildPostRequestBody builds the body of the POST request to OpenAI.
func buildPostRequestBody(card, theme, color1 string) ([]byte, error) {
	if card == "" || theme == "" || color1 == "" {
		return nil, fmt.Errorf("card, theme, and color are required")
	}

	prompt := fmt.Sprintf(
		"Design a detailed illustration for the tarot card '%s', in a '%s' style. Use hues of '%s'. Incorporate traditional tarot symbolism and ensure the design is vertically oriented without numbers or letters.",
		card, theme, color1,
	)

	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt": prompt,
		"model":  "dall-e-3", // Ensure you're using the correct model name
		"n":      1,
		"size":   "1024x1792",
	})
	if err != nil {
		return nil, err
	}

	return requestBody, nil
}

// createRequest creates an HTTP POST request to the OpenAI API.
func createRequest(url, apiKey string, requestBody []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// sendRequest sends the HTTP request and reads the response.
func sendRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second, // Set a 30-second timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Ensure the response status is OK
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	fmt.Println("Response Body: ", string(body))
	return body, nil
}

// parseImageResponse parses the OpenAI API response to extract the image URL.
func parseImageResponse(responseBody []byte) (string, error) {
	var response map[string]interface{}
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		return "", fmt.Errorf("error parsing response JSON: %w", err)
	}

	data, ok := response["data"].([]interface{})
	if !ok || len(data) == 0 {
		return "", fmt.Errorf("no images returned in the response")
	}

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
