package main

type video struct {
	Id          string
	Title       string
	Description string
	Imageurl    string
}

func getVideos() (videos []video) {

	v1 := video{
		Id:          "23223",
		Title:       "Hello Image",
		Description: "Very Nice",
		Imageurl:    "http://goodday.com",
	}
	v2 := video{
		Id:          "34646",
		Title:       "Good Image",
		Description: "Nice View",
		Imageurl:    "http://baybay.com",
	}
	videos = append(videos, v1)
	videos = append(videos, v2)
	return videos

	/// return []video{v1, v2}  # can be written like this
}
