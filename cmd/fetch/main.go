package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	method  string
	headers string
	url     string
	data    string
)

func main() {

	flag.StringVar(&method, "X", "GET", "HTTP method (e.g., GET, POST, PUT, DELETE)")
	flag.StringVar(&headers, "H", "", "Headers for the request, separated by semicolons")
	flag.StringVar(&url, "url", "", "the url to visiit eg:https://google.com")
	flag.StringVar(&data, "d", "", "data to send across.")

	flag.Parse()
	var formattedHeaders HeaderStruct
	if headers != "" {
		headers, err := parseHeaders(headers)
		if err != nil {
			fmt.Println("parsing error occured")
			os.Exit(1)
		}
		formattedHeaders = headers
	}
	fetch(method, url, formattedHeaders, nil)

}

func fetch(method string, url string, headers HeaderStruct, body []byte) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("an error occured baby")
	}
	req.Header.Set("Connection", "keep-alive")
	// if headers.populated {
	// 	//do default headers population here and addon headers

	// }
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic("it is looking grim")
	}
	defer res.Body.Close()
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Printf("%s\n", string(responseBody))

}
