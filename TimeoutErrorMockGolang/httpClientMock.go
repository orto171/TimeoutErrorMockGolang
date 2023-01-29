package main

import (
	"bytes"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
)

var (
	SuccessResponse = &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString("SUCCESS"))}
)

type HttpClientMock struct {
	mock.Mock
	DoFuncMock                func(req *http.Request) (*http.Response, error)
	numberOfRetriesSoFar      int
	MaxRetriesToReturnSuccess int
}

func (m *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	m.Called()
	m.numberOfRetriesSoFar++
	if m.numberOfRetriesSoFar >= m.MaxRetriesToReturnSuccess {
		return SuccessResponse, nil
	}
	return m.DoFuncMock(req)
}

func (m *HttpClientMock) CreateDoMock(response *http.Response, err error) {
	m.DoFuncMock = func(*http.Request) (*http.Response, error) {
		return response, err
	}
}
