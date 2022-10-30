package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Profile struct {
	SlackUserName string `json:"slackUserName"`
	Age           int64  `json:"age"`
	Backend       bool   `json:"backend"`
	Bio           string `json:"bio"`
}

func home(w http.ResponseWriter, r *http.Request) {
	age, _ := strconv.ParseInt(os.Getenv("AGE"), 10, 64)
	slackUsername, bio :=
		os.Getenv("SLACK_USERNAME"),
		os.Getenv("BIO")
	details := Profile{
		SlackUserName: slackUsername,
		Age:           age,
		Backend:       true,
		Bio:           bio,
	}

	profileJson, err := json.Marshal(details)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(profileJson)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Serving on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
