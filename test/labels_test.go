package test

import (
	"encoding/json"
	"fmt"
	"github.com/lotos2512/bamboo"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testLabel      = "prod"
	resultLabelKey = "TEST-TEST-1"
)

func TestAddLabel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(addLabelStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	label := &bamboo.Label{
		Name:      testLabel,
		ResultKey: resultLabelKey,
	}

	success, resp, err := client.Labels.AddLabel(label)
	if err != nil {
		t.Error(err)
	}

	if success == false || resp.StatusCode != 204 {
		t.Error(fmt.Sprintf("Adding Label \"%s\" was unsuccessful. Returned %s", testLabel, resp.Status))
	}
}

func addLabelStub(w http.ResponseWriter, r *http.Request) {
	label := &bamboo.Label{}
	expectedURI := fmt.Sprintf("/rest/api/latest/result/%s/label.json", resultLabelKey)

	json.NewDecoder(r.Body).Decode(label)

	if label.Name != testLabel {
		http.Error(w, "Labels do not match", http.StatusBadRequest)
		return
	}

	if r.RequestURI != expectedURI {
		http.Error(w, fmt.Sprintf("URI did not match expected %s", resultLabelKey), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
