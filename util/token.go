package util

import (
	"errors"
	"net/http"
	"strings"
)

func GetToken(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authorization token")
	}
	if vals[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	return vals[1], nil
}
