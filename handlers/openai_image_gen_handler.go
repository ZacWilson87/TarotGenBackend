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

type ImageResponse struct {
	Url         string
	Description string
}

// GenerateTarotCardDesign generates a tarot card based on the given card, theme, and color using the OpenAI API.
// It returns the image URL of the generated card.
func GenerateTarotCardDesign(card, theme, color1 string) (ImageResponse, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	imageResponse := ImageResponse{}
	if apiKey == "" {
		return imageResponse, fmt.Errorf("API key not provided")
	}

	url := os.Getenv("OPENAI_BASE_URL")
	if url == "" {
		return imageResponse, fmt.Errorf("API base URL not provided")
	}

	// Build the POST request body
	requestBody, err := buildPostRequestBody(card, theme, color1)
	if err != nil {
		return imageResponse, fmt.Errorf("error building request body: %w", err)
	}

	// Create the request
	req, err := createRequest(url, apiKey, requestBody)
	if err != nil {
		return imageResponse, fmt.Errorf("error creating request: %w", err)
	}

	// Send the request and handle the response
	responseBody, err := sendRequest(req)
	if err != nil {
		return imageResponse, fmt.Errorf("error sending request to API endpoint: %w", err)
	}

	// Parse the image URL from the response
	ImageResponse, err := parseImageResponse(responseBody)
	if err != nil {
		return imageResponse, fmt.Errorf("error parsing image response: %w", err)
	}

	return ImageResponse, nil
}

// buildPostRequestBody builds the body of the POST request to OpenAI.
func buildPostRequestBody(cardDesc, theme, color1 string) ([]byte, error) {
	if cardDesc == "" || theme == "" || color1 == "" {
		return nil, fmt.Errorf("card, theme, and color are required")
	}

	prompt := fmt.Sprintf(
		`Create a detailed illustration in portrait orientation (vertical) with dimensions 
		897 x 1497 pixels, filling the entire canvas with design. The illustration should be in the 
		style of '%s' and emphasize colors of '%s'. The subject of the illustration is: '%s'.`,
		theme, color1, cardDesc,
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

func parseImageResponse(responseBody []byte) (ImageResponse, error) {
	imageObject := ImageResponse{}
	// todo: make return obect {description: string, url: string}
	var response map[string]interface{}
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		return imageObject, fmt.Errorf("error parsing response JSON: %w", err)
	}

	data, ok := response["data"].([]interface{})
	if !ok || len(data) == 0 {
		return imageObject, fmt.Errorf("no data found in the response")
	}

	firstImage, ok := data[0].(map[string]interface{})
	if !ok {
		return imageObject, fmt.Errorf("invalid image data")
	}

	imageURL, ok := firstImage["url"].(string)
	if !ok {
		return imageObject, fmt.Errorf("no URL found in the response")
	}

	imageDesc, ok := firstImage["revised_prompt"].(string)
	if !ok {
		return imageObject, fmt.Errorf("no description found in the response")
	}

	imageObject.Description = imageDesc
	imageObject.Url = imageURL

	return imageObject, nil
}
