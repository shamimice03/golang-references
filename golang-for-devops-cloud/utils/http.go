package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HttpGet(url string) {

	response, err := http.Get(url)

	if err != nil {
		log.Fatal((err))
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Status Code %d\nBody: %s\n", response.StatusCode, body)

}
