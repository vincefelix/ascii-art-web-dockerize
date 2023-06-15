package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		path       string
		want       string
		statusCode int
	}{
		{
			name:       "find /",
			method:     http.MethodGet,
			want:       "200",
			path:       "/",
			statusCode: http.StatusOK,
		},
		{
			name:       "post /",
			method:     http.MethodPost,
			path:       "/",
			want:       "405",
			statusCode: http.StatusMethodNotAllowed,
		},

		{name: "not /",
			method:     http.MethodGet,
			path:       "/index",
			want:       "404",
			statusCode: http.StatusNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, tc.path, nil)
			responseRecorder := httptest.NewRecorder()

			indexHandler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

		})
	}
}
