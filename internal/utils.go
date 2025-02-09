package utils

import (
	"errors"
	"strings"

	"github.com/google/shlex"
)

type ParsedCurlValues struct {
	Method  string
	Url     string
	Headers []string
	Body    string
}

func ParseCurl(req string) (ParsedCurlValues, error) {
	parsedReq := ParsedCurlValues{
		Method:  "GET",
		Headers: []string{},
	}

	substrs, err := shlex.Split(req)
	if err != nil {
		return parsedReq, errors.New("Could not parse Curl request")
	}

	total := len(substrs)
	for i := 0; i < total; i++ {
		switch {
		case substrs[i] == "-X" || substrs[i] == "--request":
			if i+1 < total {
				parsedReq.Method = substrs[i+1]
				i++
			} else {
				return parsedReq, errors.New("Invalid or missing Method")
			}

		case substrs[i] == "-H" || substrs[i] == "--header":
			if i+1 < total {
				parsedReq.Headers = append(parsedReq.Headers, substrs[i+1])
				i++
			} else {
				return parsedReq, errors.New("Invalid or missing Header")
			}
		case strings.HasPrefix(substrs[i], "--data"):
			if i+1 < total {
				parsedReq.Body = substrs[i+1]
			} else {
				return parsedReq, errors.New("Invalid or missing Body")
			}
		case strings.HasPrefix(substrs[i], "http://") || strings.HasPrefix(substrs[i], "https://"):
			parsedReq.Url = substrs[i]
		}
	}

	return parsedReq, nil
}
