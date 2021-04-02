package bamboo

import (
	"fmt"
	"net/http"
	"strconv"
)

// PlanService handles communication with the plan related methods
type PlanService service

type IPlanService interface {
	CreatePlanBranch(planKey, branchName string, options *PlanCreateBranchOptions) (result bool, err error)
	NumberOfPlans() (number int, err error)
	ListPlans() (plans []*Plan, err error)
	ListPlanKeys() (keys []string, err error)
	ListPlanNames() (names []string, err error)
	PlanNameMap() (map[string]string, error)
	DisablePlan(planKey string) (err error)
	DeletePlan(planKey string) (err error)
}

// PlanCreateBranchOptions specifies the optional parameters
// for the CreatePlanBranch method
type PlanCreateBranchOptions struct {
	VCSBranch string
}

// PlanResponse encapsultes a response from the plan service
type PlanResponse struct {
	*ResourceMetadata
	Plans *Plans `json:"plans"`
}

// Plans is a collection of Plan objects
type Plans struct {
	*CollectionMetadata
	PlanList []*Plan `json:"plan"`
}

// Plan is the definition of a single plan
type Plan struct {
	ShortName string   `json:"shortName,omitempty"`
	ShortKey  string   `json:"shortKey,omitempty"`
	Type      string   `json:"type,omitempty"`
	Enabled   bool     `json:"enabled,omitempty"`
	Link      *Link    `json:"link,omitempty"`
	Key       string   `json:"key,omitempty"`
	Name      string   `json:"name,omitempty"`
	PlanKey   *PlanKey `json:"planKey,omitempty"`
}

// PlanKey holds the plan-key for a plan
type PlanKey struct {
	Key string `json:"key,omitempty"`
}

// CreatePlanBranch will create a plan branch with the given branch name for the specified build
func (p *PlanService) CreatePlanBranch(planKey, branchName string, options *PlanCreateBranchOptions) (result bool, err error) {
	var u string
	if !emptyStrings(planKey, branchName) {
		u = fmt.Sprintf("plan/%s/branch/%s.json", planKey, branchName)
	} else {
		return false, &simpleError{"Project key and/or branch name cannot be empty"}
	}

	request, err := p.client.NewRequest(http.MethodPut, u, nil)
	if err != nil {
		return
	}

	if options != nil && options.VCSBranch != "" {
		values := request.URL.Query()
		values.Add("vcsBranch", options.VCSBranch)
		request.URL.RawQuery = values.Encode()
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return
	}

	if !(response.StatusCode == 200) {
		return false, &simpleError{fmt.Sprintf("Create returned %d", response.StatusCode)}
	}

	return true, nil
}

// NumberOfPlans returns the number of plans on the Bamboo server
func (p *PlanService) NumberOfPlans() (number int, err error) {
	request, err := p.client.NewRequest(http.MethodGet, "plan.json", nil)
	if err != nil {
		return
	}

	// Restrict results to one for speed
	values := request.URL.Query()
	values.Add("max-results", "1")
	request.URL.RawQuery = values.Encode()

	planResp := PlanResponse{}
	response, err := p.client.Do(request, &planResp)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		return 0, &simpleError{fmt.Sprintf("Getting the number of plans returned %s", response.Status)}
	}

	return planResp.Plans.Size, err
}

// ListPlans gets information on all plans
func (p *PlanService) ListPlans() (plans []*Plan, err error) {
	// Get number of plans to set max-results
	numPlans, err := p.NumberOfPlans()
	if err != nil {
		return
	}

	request, err := p.client.NewRequest(http.MethodGet, "plan.json", nil)
	if err != nil {
		return
	}

	q := request.URL.Query()
	q.Add("max-results", strconv.Itoa(numPlans))
	request.URL.RawQuery = q.Encode()

	planResp := PlanResponse{}
	response, err := p.client.Do(request, &planResp)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		return nil, &simpleError{fmt.Sprintf("Getting plan information returned %s", response.Status)}
	}

	return planResp.Plans.PlanList, nil
}

// ListPlanKeys get all the plan keys for all build plans on Bamboo
func (p *PlanService) ListPlanKeys() (keys []string, err error) {
	plans, err := p.ListPlans()
	if err != nil {
		return nil, err
	}
	keys = make([]string, len(plans))
	for i, p := range plans {
		keys[i] = p.Key
	}
	return keys, nil
}

// ListPlanNames returns a list of ShortNames of all plans
func (p *PlanService) ListPlanNames() (names []string, err error) {
	plans, err := p.ListPlans()
	if err != nil {
		return nil, err
	}
	names = make([]string, len(plans))

	for i, p := range plans {
		names[i] = p.ShortName
	}
	return names, nil
}

// PlanNameMap returns a map[string]string where the PlanKey is the key and the ShortName is the value
func (p *PlanService) PlanNameMap() (map[string]string, error) {
	plans, err := p.ListPlans()
	if err != nil {
		return nil, err
	}

	planMap := make(map[string]string, len(plans))

	for _, p := range plans {
		planMap[p.Key] = p.ShortName
	}
	return planMap, nil
}

// DisablePlan will disable a plan or plan branch
func (p *PlanService) DisablePlan(planKey string) (err error) {
	u := fmt.Sprintf("plan/%s/enable", planKey)
	request, err := p.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		return newRespErr(response, "Error disable plan")
	}

	return
}

func (p *PlanService) DeletePlan(planKey string) (err error) {
	u := fmt.Sprintf("plan/%s", planKey)
	request, err := p.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusNoContent {
		return newRespErr(response, "Error delete plan")
	}

	return
}
