package main

import "news.com/events/breaking"

func main() {
	todaysNews()            // function calling from same package (news.go)
	breaking.BreakingNews() // function calling from different package (breaking)
}
