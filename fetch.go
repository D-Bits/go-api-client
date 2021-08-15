package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type JSONdata struct {
	name string `json:"name"`
	job  string `json:"job"`
	// Do they belong to the Stonecutters?
	stonecutter bool `json:"stonecutter"`
}

func getData(endpoint string) {

	// make a GET request to fetch the JSON data
	res, err := http.Get(endpoint)

	// Log any errors
	if err != nil {
		log.Fatal(err)
	}
	// store the response body in a variable
	body, err := ioutil.ReadAll(res.Body)

	// Log any errors
	if err != nil {
		log.Fatal(err)
	}

	// Create a file to write to
	file, err := os.Create("./data/people.json")

	file.WriteString(string(body))

	// Log any errors
	if err != nil {
		log.Fatal(err)
	}

}
