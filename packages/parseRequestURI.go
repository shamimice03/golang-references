package main

import (
	"fmt"
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

	http.Get("https://www.amazon.com/")

}
