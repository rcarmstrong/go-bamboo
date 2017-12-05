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

// PlanResponce encapsultes a response from the REST API
type PlanResponce struct {
	Response Plans `json:"plans"`
}

// Plans is a collection of Plan objects
type Plans struct {
	Size int    `json:"size"`
	List []Plan `json:"plan"`
}

// Plan is the definition of a single plan
type Plan struct {
	ShortName string  `json:"shortName"`
	ShortKey  string  `json:"shortKey"`
	Type      string  `json:"type"`
	Enabled   bool    `json:"enabled"`
	Key       string  `json:"key"`
	Name      string  `json:"name"`
	PK        PlanKey `json:"planKey"`
}

// PlanKey holds the plan-key for a plan
type PlanKey struct {
	Key string `json:"key"`
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

	planInfo := PlanResponce{}
	resp, err := p.client.Do(req, &planInfo)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, &simpleError{fmt.Sprintf("Getting the number of plans returned %s", resp.Status)}
	}

	return planInfo.Response.Size, nil
}

// ListPlans gets information on all plans
func (p *PlanService) ListPlans() ([]Plan, error) {

	numPlans, err := p.NumberOfPlans()
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("plan.json?max-results=%d", numPlans)

	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	planInfo := PlanResponce{}
	resp, err := p.client.Do(req, &planInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, &simpleError{fmt.Sprintf("Getting plan information returned %s", resp.Status)}
	}

	return planInfo.Response.List, nil
}
