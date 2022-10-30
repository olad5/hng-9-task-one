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
	SlackUserName string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int64  `json:"age"`
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
  port := os.Getenv("PORT")
	http.HandleFunc("/", home)
	fmt.Println("Serving on :" + port)
  log.Fatal(http.ListenAndServe(":"+ port, nil))
}
