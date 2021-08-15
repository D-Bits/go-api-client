package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getData(endpoint string, filename string) {

	// make a GET request to fetch the JSON data
	res, err := http.Get(endpoint)

	// Log any errors
	if err != nil {
		log.Fatal(err)
	}

	// store the response body in a variable
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Create a file to write to
	file, err := os.Create("./data/" + filename + ".json")
	// Write the JSON to an external file
	file.WriteString(string(body))

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Download complete. Check the 'data' directory.")
	}

}
