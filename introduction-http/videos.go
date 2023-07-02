package main

import (
	"encoding/json" // Package for JSON encoding and decoding
	"io/ioutil"     // Package for file I/O operations
)

// Struct representing the attributes of a video
type video struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Imageurl    string `json:"imageurl"`
	Url         string `json:"url"`
}

// Function to get videos from a JSON file
func getVideos() (videos []video) {

	filebyets, err := ioutil.ReadFile("./videos.json") // Read file content as byte slice

	if err != nil {
		panic(err) // Panic if there's an error reading the file
	}

	err = json.Unmarshal(filebyets, &videos) // Unmarshal JSON data into videos slice

	if err != nil {
		panic(err) // Panic if there's an error unmarshaling the JSON data
	}

	return videos // Return the populated videos slice
}

// Function to save videos to a JSON file
func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos) // Marshal videos slice into JSON format

	if err != nil {
		panic(err) // Panic if there's an error marshaling the JSON data
	}

	ioutil.WriteFile("./videos-updated.json", videoBytes, 0644) // Write JSON data to file
}
