package bamboo_test

import (
	"os"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

var (
	bambooCLI *bamboo.Client
)

func init() {
	b := bamboo.NewSimpleClient(nil, os.Getenv("BAMBOO_USERNAME"), os.Getenv("BAMBOO_PASSWORD"))
	bambooCLI = b
}

func TestNumberOfPlans(t *testing.T) {
	numPlans, err := bambooCLI.Plans.NumberOfPlans()
	if err != nil {
		t.Error(err)
	}
	t.Log(numPlans)
}

func TestListPlanKeys(t *testing.T) {
	t.Log(bambooCLI.BaseURL.String())
	keys, err := bambooCLI.Plans.ListPlanKeys()
	if err != nil {
		t.Errorf("Error listing plan keys %s", err)
	}
	t.Log(keys)
}
