package main

import (
	"github.com/shamimice03/golang-references/golang-for-devops-cloud/utils"
)

// Input
// go run . http://test.com

func main() {

	println("Args url validation")
	utils.Args_url_validation()

	println("Http Get")
	url := "https://www.amazon.com/"
	utils.HttpGet(url)

	println("Url Details")
	urlString := "https://example.com/path?param1=value1&param2=value2"
	utils.UrlDetails(urlString)
}
