package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

var (
	httpClientMock = &HttpClientMock{}
)

func arrangeTestMock() {
	timeoutError := CreateTimeoutErrorMock()
	expectedResponse := &http.Response{StatusCode: http.StatusRequestTimeout,
		Body: io.NopCloser(bytes.NewBufferString("TIMEOUT ERROR"))}
	httpClientMock.CreateDoMock(expectedResponse, timeoutError)
	httpClientMock.On("Do").Return(expectedResponse, timeoutError)
}

func TestHttpClientTimeoutError(t *testing.T) {

	// arrange
	arrangeTestMock()
	httpClientMock.MaxRetriesToReturnSuccess = 5

	// act
	err := sendRequest(httpClientMock)

	// assert
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "TIMEOUT ERROR")

}

func TestHttpClientSucceededAfterRetry(t *testing.T) {

	// arrange
	arrangeTestMock()
	httpClientMock.MaxRetriesToReturnSuccess = 2

	// act
	err := sendRequest(httpClientMock)

	// assert
	assert.NoError(t, err)

}
