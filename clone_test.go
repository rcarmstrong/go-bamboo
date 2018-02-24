package bamboo_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

func TestClonePlan(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(clonePlanStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	planClone, resp, err := client.Clone.ClonePlan("CORE-TEST", "CORE-TESTS")
	if err != nil {
		t.Error(err)
	}

	if planClone.Key != "CORE-TESTS" || resp.StatusCode != 200 {
		t.Error(fmt.Sprintf("Creating plan clone \"CORE-TESTS\" was unsuccessful. Returned %s", resp.Status))
	}
}

func clonePlanStub(w http.ResponseWriter, r *http.Request) {

	planClone := bamboo.Plan{
		Key: "CORE-TESTS",
	}

	bytes, err := json.Marshal(planClone)
	if err != nil {
		panic(err)
	}

	w.Write(bytes)

}
