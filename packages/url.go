package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "https://example.com/path?param1=value1&param2=value2"
	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	// Access various components of the parsed URL
	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("Query:", u.Query())
}

// Scheme: https
// Host: example.com
// Path: /path
// Query: map[param1:[value1] param2:[value2]]
