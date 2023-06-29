package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	UserInfo Info   `json:"userinfo"`
}

type Info struct {
	UserId     string `json:"userid"`
	Department string `json:"department"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Hey, Hello</h1>"))
}

func GetBio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	user := User{
		Name: "John",
		Age:  "28",
		UserInfo: Info{
			UserId:     "656321",
			Department: "HR",
		},
	}

	json.NewEncoder(w).Encode(user)
}

func main() {

	http.HandleFunc("/greetings", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hey, Good Day!"))
	})

	http.HandleFunc("/sayhello", SayHello)
	http.HandleFunc("/getbio", GetBio)

	fmt.Println("Staring server at port 6000")
	log.Fatal(http.ListenAndServe(":6000", nil))

}

//iR6fEsj5JaTo
