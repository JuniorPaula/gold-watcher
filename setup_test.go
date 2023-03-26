package main

import (
	"bytes"
	"goldwatcher/repository"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `
	{
		"ts": 1679610041155,
		"tsj": 1679610035453,
		"date": "Mar 23rd 2023, 06:20:35 pm NY",
		"items": [
		{
			"curr": "BRL",
			"xauPrice": 10557.461,
			"xagPrice": 122.3563,
			"chgXau": 314.5546,
			"chgXag": 3.1337,
			"pcXau": 3.071,
			"pcXag": 2.6284,
			"xauClose": 10242.90644,
			"xagClose": 119.22258
		}
		]
	}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
