package main

import "fmt"

func main() {
	videos := getVideos() // Get videos from the JSON file

	// Iterate over each video in the videos slice
	for k, _ := range videos {
		videos[k].Description = videos[k].Description + "\n Adding Demo Description" // Append additional text to the video's description
	}

	fmt.Println(videos) // Print the modified videos slice

	saveVideos(videos) // Save the updated videos slice to a JSON file
}
