package bamboo_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

func TestLatestResult(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(latestResultStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, resp, err := client.Results.LatestResult("CORE-TEST")
	if err != nil {
		if resp != nil {
			t.Log(resp.Status)
		}
		t.Error(err)
	}
}

func latestResultStub(w http.ResponseWriter, r *http.Request) {
	var expectedURL = "CORE-TEST-latest"

	check := strings.Split(strings.Split(r.URL.String(), "result/")[1], "?")[0]
	if check == expectedURL && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestNumberedResult(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(numberedResultStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, resp, err := client.Results.NumberedResult("CORE-TEST-1")
	if err != nil {
		if resp != nil {
			t.Log(resp.Status)
		}
		t.Error(err)
	}
}

func numberedResultStub(w http.ResponseWriter, r *http.Request) {
	var expectedURL = "CORE-TEST-1"

	check := strings.Split(strings.Split(r.URL.String(), "result/")[1], "?")[0]
	if check == expectedURL && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
