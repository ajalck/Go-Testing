package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name        string
	server      *httptest.Server
	expResponse *Weather
	expError    error
}

func TestGetWeather(t *testing.T) {
	test := []Tests{
		{
			name: "basic-request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"city":"Kochi","forecast":"Sunny"}`))
			})),
			expResponse: &Weather{
				City:     "Kochi",
				Forecast: "Sunny",
			},
			expError: nil,
		},
	}
	for _, tst := range test {
		t.Run(tst.name, func(t *testing.T) {
			defer tst.server.Close()

			resp, err := GetWeather(tst.server.URL)

			if !reflect.DeepEqual(resp, tst.expResponse) {
				t.Errorf("FAILED : expected:%v ,got:%v\n", tst.expResponse, resp)
			}
			if !errors.Is(err, tst.expError) {
				t.Errorf("Expected error FAILED : expected:%v ,got:%v\n", tst.expError, err)
			}
		})
	}
}
