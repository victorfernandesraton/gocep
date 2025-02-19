package handler

import (
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
)

// go test -run ^TestSearchCep$ -v
func TestSearchCep(t *testing.T) {
	type args struct {
		method string
		ctype  string
		Header map[string]string
		url    string
	}
	tests := []struct {
		name     string
		args     args
		want     int //statuscode
		bodyShow bool
	}{
		// [GET] /v1/cep/xxxx
		{"test_searchcep_", args{"GET", "application/json", nil, "/v1/cep/0"}, 400, false},
		{"test_searchcep_", args{"GET", "application/json", nil, "/v1/cep/08226-021"}, 400, false},
		{"test_searchcep_", args{"GET", "application/json", nil, "/v1/cep/08226021"}, 200, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tt.args.method, tt.args.url, nil)
			req.Header.Add("Content-type", tt.args.ctype)
			for key, val := range tt.args.Header {
				req.Header.Add(key, val)
			}
			SearchCep(w, req)
			resp := w.Result()
			defer resp.Body.Close()

			if !reflect.DeepEqual(resp.StatusCode, tt.want) {
				t.Errorf("Call() out status = %v, want status %v", resp.StatusCode, tt.want)
				return
			}
			// assert.Equal(t, tt.want, resp.StatusCode)
			if tt.bodyShow {
				body, _ := io.ReadAll(resp.Body)
				t.Log("\n Resp : \n", string(body), "\n")
			}
		})
	}
}
