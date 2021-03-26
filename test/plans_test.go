package test

import (
	"github.com/lotos2512/bamboo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
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
