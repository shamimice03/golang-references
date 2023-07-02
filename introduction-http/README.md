### v1 (`main.go`)
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil) // (nil) means working with default multiplexer
}

// Hello is the HTTP request handler function for the root path ("/").
func Hello(w http.ResponseWriter, r *http.Request) {

	// Iterate through the request headers and print the key-value pairs.
	for key, value := range r.Header {
		fmt.Printf("Key: %v\t Value: %v\n", key, value)
	}

	// Add a custom header to the response.
	w.Header().Add("Test", "HelloFromResponseHeader")

	// Write the "hello" string to the response.
	w.Write([]byte("hello"))
}
```

### v2 (`main.go`)
```go
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
	http.ListenAndServe(":8080", nil) // (nil) means working with default multiplexer
}

// HandleGetVideos handles the GET request to retrieve videos
func HandleGetVideos(w http.ResponseWriter, r *http.Request) {

	vidoes := getVideos() // Retrieve videos
	videoBytes, err := json.Marshal(vidoes) // Convert videos to JSON

	if err != nil {
		panic(nil) // Panic if there's an error in JSON marshaling
	}

	w.Write(videoBytes) // Write JSON response
}

// HandleUpdateVideos handles the POST request to update videos
func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body) // Read request body

		if err != nil {
			panic(err) // Panic if there's an error reading the request body
		}

		var videos []video
		err = json.Unmarshal(body, &videos) // Convert request body (JSON) to videos

		if err != nil {
			w.WriteHeader(400) // Set response status code to 400 (Bad Request)
			fmt.Fprint(w, "Bad Request") // Write response message
		}

		saveVideos(videos) // Save the updated videos

	} else {
		w.WriteHeader(405) // Set response status code to 405 (Method Not Allowed)
		fmt.Fprint(w, "Method not supported") // Write response message
	}
}

```

### Explanation:
This code is a simple Go program that sets up an HTTP server. It defines two request handlers, `HandleGetVideos` and `HandleUpdateVideos`, for handling GET and POST requests, respectively.

The `HandleGetVideos` function retrieves videos, converts them to JSON using `json.Marshal`, and writes the JSON response to the `http.ResponseWriter`.

The `HandleUpdateVideos` function checks if the request method is POST. If it is, it reads the request body using `ioutil.ReadAll` and attempts to unmarshal the JSON data into a slice of `video` structs. If there's an error in unmarshaling, it sets the response status code to 400 (Bad Request) and writes a response message.

If the request method is not POST, it sets the response status code to 405 (Method Not Allowed) and writes a response message.

Both functions are registered as handlers for specific paths using `http.HandleFunc`. Finally, `http.ListenAndServe` starts the server on port 8080 using the default multiplexer.

### Postman
![Alt text](<Screenshot 2023-07-03 at 2.02.14.png>)

### Ref:
- https://www.youtube.com/watch?v=MKkokYpGyTU&t=138s&ab_channel=ThatDevOpsGuy
- https://github.com/marcel-dempers/docker-development-youtube-series/tree/master/golang/introduction/part-3.http

