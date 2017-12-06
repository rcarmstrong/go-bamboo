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
	ServiceMetadata
	Plans Plans `json:"plans"`
}

// Plans is a collection of Plan objects
type Plans struct {
	CollectionMetadata
	PlanList []Plan `json:"plan"`
}

// Plan is the definition of a single plan
type Plan struct {
	ShortName string      `json:"shortName,omitempty"`
	ShortKey  string      `json:"shortKey,omitempty"`
	Type      string      `json:"type,omitempty"`
	Enabled   bool        `json:"enabled,omitempty"`
	Link      ServiceLink `json:"link,omitempty"`
	Key       string      `json:"key,omitempty"`
	Name      string      `json:"name,omitempty"`
	PlanKey   PlanKey     `json:"planKey,omitempty"`
}

// PlanKey holds the plan-key for a plan
type PlanKey struct {
	Key string `json:"key,omitempty"`
}

// CreatePlanBranch will create a plan branch with the given branch name for the specified build
func (p *PlanService) CreatePlanBranch(projectKey, buildKey, branchName string, opt *PlanCreateBranchOptions) (*http.Response, error) {
	var u string
	if !emptyStrings(projectKey, buildKey, branchName) {
		u = fmt.Sprintf("plan/%s-%s/branch/%s.json", projectKey, buildKey, branchName)
		if opt != nil && opt.VCSBranch != "" {
			u += fmt.Sprintf("?vcsBranch=%s", opt.VCSBranch)
		}
	} else {
		return nil, &simpleError{"Project key, build key, and branch name cannot be empty"}
	}

	req, err := p.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == 200) {
		return resp, &simpleError{fmt.Sprintf("Create returned %d", resp.StatusCode)}
	}

	return resp, nil
}

// NumberOfPlans returns the number of plans on the Bamboo server
func (p *PlanService) NumberOfPlans() (int, error) {
	// Restrict results to one for speed
	u := "plan.json?max-results=1"

	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return 0, err
	}

	planResp := PlanResponse{}
	resp, err := p.client.Do(req, &planResp)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, &simpleError{fmt.Sprintf("Getting the number of plans returned %s", resp.Status)}
	}

	return planResp.Plans.Size, nil
}

// ListPlans gets information on all plans
func (p *PlanService) ListPlans() (*Plans, error) {

	numPlans, err := p.NumberOfPlans()
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("plan.json?max-results=%d", numPlans)

	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	planResp := PlanResponse{}
	resp, err := p.client.Do(req, &planResp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, &simpleError{fmt.Sprintf("Getting plan information returned %s", resp.Status)}
	}

	return &planResp.Plans, nil
}

// ListPlanKeys get all the plan keys for all build plans on Bamboo
func (p *PlanService) ListPlanKeys() ([]string, error) {
	plans, err := p.ListPlans()
	if err != nil {
		return nil, err
	}
	keys := make([]string, plans.Size)

	for _, p := range plans.PlanList {
		keys = append(keys, p.PlanKey.Key)
	}
	return keys, nil
}
