package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Color struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Year         int    `json:"year"`
	Color        string `json:"color"`
	PantoneValue string `json:"pantone_value"`
}

type Support struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

type Response struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       []Color `json:"data"`
	Support    Support `json:"support"`
}

func main() {
	url := "https://reqres.in/api/unknown"
	response, err := makeRequest(url)
	if err != nil {
		log.Fatal(err)
	}

	printResponse(response)
	accessResponse(response)

}

// makeRequest performs an HTTP GET request and returns the response as a *Response or an error if any.
func makeRequest(url string) (*Response, error) {
	// Create an HTTP client with a timeout of 10 seconds
	client := &http.Client{Timeout: 10 * time.Second}

	// Send the GET request
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the response body into a Response struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return &response, nil
}

// printResponse prints the response data.
func printResponse(response *Response) {
	// Marshal the response to JSON with indentation
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal response to JSON: %v", err)
	}

	// Print the formatted JSON and the length of the Data array
	fmt.Println(string(jsonData))
	fmt.Println("Data length:", len(response.Data))
}

// accessResponse iterates over the response data and prints the key-value pairs.
func accessResponse(response *Response) {
	for key, value := range response.Data {
		fmt.Println(key, ": ", value)
	}

	for _, value := range response.Data {
		fmt.Println(value.Name, ": ", value.Color)
	}
}
