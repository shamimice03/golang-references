package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	// Command line arguments
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Enter An URL")
		os.Exit(1)
	}

	if url, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Println(url, err)
		os.Exit(1) // Invalid
	}

	response, err := http.Get("https://www.amazon.com/")

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
