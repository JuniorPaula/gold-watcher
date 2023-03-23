package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {

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
