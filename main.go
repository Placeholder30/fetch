package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	URL "net/url"
	"os"
)

var (
	method   string
	headers  string
	url      string
	data     string
	location string
)

func main() {

	flag.StringVar(&method, "X", "GET", "HTTP method (e.g., GET, POST, PUT, DELETE)")
	flag.StringVar(&headers, "H", "", "Headers for the request, separated by semicolons")
	flag.StringVar(&url, "url", "", "the url to visit eg:https://google.com")
	flag.StringVar(&data, "d", "", "data to send across")
	flag.StringVar(&location, "o", "", "out location")
	flag.Parse()

	if _, err := URL.ParseRequestURI(url); err != nil {
		fmt.Println("invalid url", url)
		os.Exit(1)
	}

	var formattedHeaders = make(map[string]string)
	if headers != "" {
		headers, err := parseHeaders(headers)
		if err != nil {
			fmt.Println("parsing error occured")
			os.Exit(1)
		}
		formattedHeaders = headers
	} else {
		formattedHeaders["Content-Type"] = "text/plain"
	}

	fetch(method, url, formattedHeaders, []byte(data), location)

}

func fetch(method string, url string, headers map[string]string, body []byte, o string) {
	var bodyData = bytes.NewBuffer(body)
	fmt.Println(method, url, bodyData)
	req, err := http.NewRequest(method, url, bodyData)
	if err != nil {
		fmt.Println("an error occured baby")
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("it is looking grim")
	}
	defer res.Body.Close()

	if len(o) > 1 {
		out, err := os.Create(o)
		if err != nil {
			fmt.Println("could not open file")
			os.Exit(1)
		}
		defer out.Close()
		n, err := io.Copy(out, res.Body)
		if err != nil {
			fmt.Println("an error occured mate, soz")
		}
		fmt.Println(n)
		return

	} else {
		_, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			return
		}
		// fmt.Printf("%s\n", string(responseBody))
	}

}
