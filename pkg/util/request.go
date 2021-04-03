package util

import (
	"io"
	"net/http"
)

func PostJSON(url string, cType string, value io.Reader) *http.Response {
	res, err := http.Post(url, cType, value)
	ErrorHandler(err, "panic")
	return res
}
