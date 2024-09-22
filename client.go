package main

import (
	"net/http"
	"log"
	"io"
	"bytes"
	"net/url"
	"encoding/json"
)

const BASE_URL = "127.0.0.1:8080"

type PostBody struct {
	Title string `json:"title"`
}

func main(){

	// Client get request
	resp2, err := http.Get("http://127.0.0.1:8080/.well-known/jwks.json")
	if err != nil {
	  log.Fatal(err)
	}
	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
	   log.Fatalln(err)
	}
	sb2 := string(body2)
   	log.Printf(sb2)

	// Client post request
	// Prepare URL
	postURL := url.URL{
		Host:   BASE_URL,
		Path:   "/auth",
		Scheme: "http",
	}

	// Prepare request body
	body := PostBody{
		Title: "How much wood could a woodchuck chuck if a woodchuck could chuck wood?",
	}

	bodyBytes, err := json.Marshal(&body)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(bodyBytes)

	// Make HTTP POST request
	resp, err := http.Post(postURL.String(), "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}

	// Close response body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
		log.Println("Error response. Status Code: ", resp.StatusCode)
	}

	log.Println(string(responseBody))
}