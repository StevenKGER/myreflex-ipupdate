package api

import (
	"bytes"
	"log"
	"net/http"
)

var httpClient = http.Client{}

func post(url string, data []byte) (response *http.Response) {
	return do("POST", url, bytes.NewBuffer(data))
}

func patch(url string, data []byte) (response *http.Response) {
	return do("PATCH", url, bytes.NewBuffer(data))
}

func get(url string) *http.Response {
	return do("GET", url, bytes.NewBuffer([]byte{}))
}

func do(method string, url string, data *bytes.Buffer) (response *http.Response) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		log.Fatalln("An error occurred while creating a request", err)
	}

	request.Header.Set("User-Agent", userAgent)

	response, err = httpClient.Do(request)
	if err != nil {
		log.Fatalln("An error occurred while sending a request", err)
	}

	return response
}
