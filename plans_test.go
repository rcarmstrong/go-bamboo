package bamboo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

func TestListPlans(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(unauthorizedStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, response, err := client.Plans.ListPlanKeys()
	assert.Error(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestListPlanNames(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(unauthorizedStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, response, err := client.Plans.ListPlanNames()
	assert.Error(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}
