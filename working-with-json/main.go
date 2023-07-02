package main

import "fmt"

func main() {

	videos := getVideos()
	for k, _ := range videos {

		videos[k].Description = videos[k].Description + "\n Adding Des"
	}
	fmt.Println(videos)
	saveVideos(videos)
}
