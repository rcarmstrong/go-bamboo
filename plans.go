package bamboo

import (
	"fmt"
	"net/http"
)

// PlanService handles communication with the plan related methods
type PlanService service

// PlanCreateBranchOptions specifies the optional parameters
// for the CreatePlanBranch method
type PlanCreateBranchOptions struct {
	VCSBranch string
}

// PlanResponse encapsultes a response from the plan service
type PlanResponse struct {
	*ServiceMetadata
	Plans *Plans `json:"plans"`
}

// Plans is a collection of Plan objects
type Plans struct {
	*CollectionMetadata
	PlanList []*Plan `json:"plan"`
}

// Plan is the definition of a single plan
type Plan struct {
	ShortName string       `json:"shortName,omitempty"`
	ShortKey  string       `json:"shortKey,omitempty"`
	Type      string       `json:"type,omitempty"`
	Enabled   bool         `json:"enabled,omitempty"`
	Link      *ServiceLink `json:"link,omitempty"`
	Key       string       `json:"key,omitempty"`
	Name      string       `json:"name,omitempty"`
	PlanKey   *PlanKey     `json:"planKey,omitempty"`
}

// PlanKey holds the plan-key for a plan
type PlanKey struct {
	Key string `json:"key,omitempty"`
}

// CreatePlanBranch will create a plan branch with the given branch name for the specified build
func (p *PlanService) CreatePlanBranch(projectKey, buildKey, branchName string, options *PlanCreateBranchOptions) (bool, *http.Response, error) {
	var u string
	if !emptyStrings(projectKey, buildKey, branchName) {
		u = fmt.Sprintf("plan/%s-%s/branch/%s.json", projectKey, buildKey, branchName)
	} else {
		return false, nil, &simpleError{"Project key, build key, and branch name cannot be empty"}
	}

	request, err := p.client.NewRequest("PUT", u, nil)
	if err != nil {
		return false, nil, err
	}

	if options != nil && options.VCSBranch != "" {
		request.URL.Query().Add("vcsBranch", options.VCSBranch)
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return false, response, err
	}

	if !(response.StatusCode == 200) {
		return false, response, &simpleError{fmt.Sprintf("Create returned %d", response.StatusCode)}
	}

	return false, response, nil
}

// NumberOfPlans returns the number of plans on the Bamboo server
func (p *PlanService) NumberOfPlans() (int, *http.Response, error) {
	request, err := p.client.NewRequest("GET", "plan.json", nil)
	if err != nil {
		return 0, nil, err
	}

	// Restrict results to one for speed
	request.URL.Query().Add("max-results", "1")

	planResp := PlanResponse{}
	response, err := p.client.Do(request, &planResp)
	if err != nil {
		return 0, response, err
	}

	if response.StatusCode != 200 {
		return 0, response, &simpleError{fmt.Sprintf("Getting the number of plans returned %s", response.Status)}
	}

	return planResp.Plans.Size, response, nil
}

// ListPlans gets information on all plans
func (p *PlanService) ListPlans() ([]*Plan, *http.Response, error) {
	numPlans, resp, err := p.NumberOfPlans()
	if err != nil {
		return nil, resp, err
	}

	request, err := p.client.NewRequest("GET", "plan.json", nil)
	if err != nil {
		return nil, nil, err
	}

	request.URL.Query().Add("max-results", string(numPlans))

	planResp := PlanResponse{}
	response, err := p.client.Do(request, &planResp)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Getting plan information returned %s", response.Status)}
	}

	return planResp.Plans.PlanList, response, nil
}

// ListPlanKeys get all the plan keys for all build plans on Bamboo
func (p *PlanService) ListPlanKeys() ([]string, *http.Response, error) {
	plans, response, err := p.ListPlans()
	if err != nil {
		return nil, response, err
	}
	keys := make([]string, len(plans))

	for i, p := range plans {
		keys[i] = p.ShortKey
	}
	return keys, response, nil
}
