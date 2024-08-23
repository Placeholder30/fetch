package main

import (
	"errors"
	"strings"
)

var headersMap = make(map[string]string)

func parseHeaders(headers string) (map[string]string, error) {
	splitHeaders := strings.Split(headers, ";")
	if len(splitHeaders) < 1 {
		return map[string]string{}, errors.New("invalid headers format ")

	}

	for _, header := range splitHeaders {
		headerSlice := strings.Split(header, ":")
		if len(headerSlice) != 2 {
			continue
		}
		headerKey := strings.TrimSpace(headerSlice[0])
		headerValue := strings.TrimSpace(headerSlice[1])
		headersMap[headerKey] = headerValue
	}
	return headersMap, nil
}
