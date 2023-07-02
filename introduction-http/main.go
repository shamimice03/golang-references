package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", HandleGetVideos)
	http.HandleFunc("/update", HandleUpdateVideos)
	http.ListenAndServe(":8080", nil) // (nil) means working with default multiplaxer
}

func HandleGetVideos(w http.ResponseWriter, r *http.Request) {

	vidoes := getVideos()
	videoBytes, err := json.Marshal(vidoes)

	if err != nil {
		panic(nil)
	}

	w.Write(videoBytes)
}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		var videos []video
		err = json.Unmarshal(body, &videos)

		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, "Bad Request")
		}

		saveVideos(videos)

	} else {
		w.WriteHeader(405)
		fmt.Fprint(w, "Method nor supported")
	}
}
