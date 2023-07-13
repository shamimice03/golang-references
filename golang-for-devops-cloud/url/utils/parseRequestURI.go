package utils

import (
	"fmt"
	"net/url"
	"os"
)

func Args_url_validation() {

	// Command line arguments
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Enter An URL")
		os.Exit(1)
	}

	if url, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Println(url, err)
		os.Exit(1) // Invalid
	} else {
		fmt.Println(url)
	}

}
