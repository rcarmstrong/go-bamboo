package bamboo_test

import (
	"testing"
)

func TestNumberOfPlans(t *testing.T) {
	numPlans, err := bambooCLI.Plans.NumberOfPlans()
	if err != nil {
		t.Error(err)
	}
	t.Log(numPlans)
}

func TestListPlanKeys(t *testing.T) {
	keys, err := bambooCLI.Plans.ListPlanKeys()
	if err != nil {
		t.Errorf("Error listing plan keys %s", err)
	}
	t.Log(keys)
}
