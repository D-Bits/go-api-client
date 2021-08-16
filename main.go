package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Create a reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a URL: ")
	url, _ := reader.ReadString('\n')
	// Remove the newline before passing into getData() function
	formattedUrl := strings.TrimSuffix(url, "\n")
	fmt.Print("Enter a file name: ")
	filename, _ := reader.ReadString('\n')
	formattedFilename := strings.TrimSuffix(filename, "\n")

	getData(formattedUrl, formattedFilename)

}
