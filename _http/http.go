package _http

import (
	"net/http"
)

func GetHttpResponseBody(url string) (response *http.Response, err error) {
	response, err = http.Get(url)
	return response, err
}
