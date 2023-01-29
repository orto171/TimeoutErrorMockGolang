package main

import (
	"errors"
	"net"
	"net/http"
)

var (
	retryNumber      = 0
	maxRetryAttempts = 3
)

func sendRequest(httpClient HTTPClient) error {
	req, err := http.NewRequest(http.MethodGet, "https://www.google.com/", nil)
	if err != nil {
		return err
	}
	resp, httpError := httpClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	} else if timeoutError, ok := httpError.(net.Error); ok && timeoutError.Timeout() {
		if retryNumber < maxRetryAttempts {
			retryNumber += 1
			return sendRequest(httpClient)
		} else {
			return errors.New("TIMEOUT ERROR")
		}
	} else {
		return err
	}
}
