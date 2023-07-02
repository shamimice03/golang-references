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

	// The following code demonstrates manually inserting data into the videos slice
	// v1 := video{
	// 	Id:          "23223",
	// 	Title:       "Hello Image",
	// 	Description: "Very Nice",
	// 	Imageurl:    "http://goodday.com",
	// 	Url:         "http://test.com",
	// }
	// v2 := video{
	// 	Id:          "34646",
	// 	Title:       "Good Image",
	// 	Description: "Nice View",
	// 	Imageurl:    "http://baybay.com",
	// 	Url:         "http://testting.com",
	// }
	// videos = append(videos, v1)
	// videos = append(videos, v2)
	// return videos
	// The above code can be written as: return []video{v1, v2}
}

// Function to save videos to a JSON file
func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos) // Marshal videos slice into JSON format

	if err != nil {
		panic(err) // Panic if there's an error marshaling the JSON data
	}

	ioutil.WriteFile("./videos-updated.json", videoBytes, 0644) // Write JSON data to file
}
