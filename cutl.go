package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

// CURLGet .
func CURLGet(reqURL string) ([]byte, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}
