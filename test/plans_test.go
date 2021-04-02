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

	_, err := client.Plans.ListPlanKeys()
	assert.Error(t, err)
}

func TestListPlanNames(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(unauthorizedStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, err := client.Plans.ListPlanNames()
	assert.Error(t, err)
}
