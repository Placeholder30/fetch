package main

import (
	"errors"
	"strings"
)

type HeaderStruct struct {
	Authorization string
	CacheControl  string `json:"Cache-Control"`
	ContentType   string `json:"Content-Type"`
	Forwarded     string
	Host          string
	Origin        string
	UserAgent     string `json:"User-Agent"`
	populated     bool
}

var approvedHeaders HeaderStruct

func parseHeaders(headers string) (HeaderStruct, error) {
	splitHeaders := strings.Split(headers, ";")
	if len(splitHeaders) < 2 {
		return HeaderStruct{}, errors.New("invalid headers format ")

	}

	for _, header := range splitHeaders {
		headerSlice := strings.Split(header, ":")
		if len(headerSlice) != 2 {
			continue
		}
		headerKey := strings.TrimSpace(headerSlice[0])
		headerValue := strings.TrimSpace(headerSlice[1])

		switch headerKey {
		case "Authorization":
			approvedHeaders.Authorization = headerValue
			approvedHeaders.populated = true
		case "Cache-Control":
			approvedHeaders.CacheControl = headerValue
			approvedHeaders.populated = true
		case "Content-Type":
			approvedHeaders.ContentType = headerValue
			approvedHeaders.populated = true
		case "Forwarded  ":
			approvedHeaders.Forwarded = headerValue
			approvedHeaders.populated = true
		case "Host ":
			approvedHeaders.Host = headerValue
			approvedHeaders.populated = true
		case "Origin":
			approvedHeaders.Origin = headerValue
			approvedHeaders.populated = true
		case "User-Agent ":
			approvedHeaders.UserAgent = headerValue
			approvedHeaders.populated = true
		}
	}
	return approvedHeaders, nil
}
