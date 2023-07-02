package main

import (
	"encoding/json"
	"io/ioutil"
)

type video struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Imageurl    string `json:"imageurl"`
	Url         string `json:"url"`
}

func getVideos() (videos []video) {

	filebyets, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(filebyets, &videos)

	if err != nil {
		panic(err)
	}

	return videos

	// convert the byte to string
	// fileContent := string(filebyets)
	// fmt.Printf(fileContent)
	// return nil

	// manually insert data to the struct (video)

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

	/// return []video{v1, v2}  # can be written like this
}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)
}
