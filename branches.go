package bamboo

import (
	"fmt"
	"net/http"
)

// PlanBranchService is a derivative of the plan service to handle
// interacting with plan branches
type PlanBranchService service

// BranchesResponse encapsulates the information from
// requesting plan branch information
type BranchesResponse struct {
	*ResourceMetadata
	Branches *Branches `json:"branches"`
}

// Branches is the collection of branches
type Branches struct {
	*CollectionMetadata
	BranchList []*Branch `json:"branch"`
}

// Branch represents a single plan branch
type Branch struct {
	Description  string `json:"description"`
	ShortName    string `json:"shortName"`
	ShortKey     string `json:"shortKey"`
	Enabled      bool   `json:"enabled"`
	Link         *Link
	WorkflowType string `json:"workflowType"`
	*PlanKey
	Name string `json:"name,omitempty"`
}

// PlanBranchExpandOptions are the optional parameters to a request
// for plan branch information.
type PlanBranchExpandOptions struct {
	Actions         bool
	Stages          bool
	Branches        bool
	VariableContext bool
}

// ListPlanBranches lists all plan branches for a given plan
func (pb *PlanBranchService) ListPlanBranches(planKey string) ([]*Branch, *http.Response, error) {
	u := fmt.Sprintf("plan/%s/.json", planKey)

	request, err := pb.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	q := request.URL.Query()
	// Setting max-results very high to try and get all branches
	q.Set("max-results", "10000")
	// Expand branch information of the given plan
	q.Set("expand", "branches")
	request.URL.RawQuery = q.Encode()

	branchResponse := BranchesResponse{}
	response, err := pb.client.Do(request, &branchResponse)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Listing plan branches for %s returned %s", planKey, response.Status)}
	}

	return branchResponse.Branches.BranchList, response, err
}

// ListVCSBranches returns a list of all VCS branches
func (pb *PlanBranchService) ListVCSBranches(planKey string) ([]string, *http.Response, error) {
	u := fmt.Sprintf("plan/%s/vcsBranches.json", planKey)

	request, err := pb.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	q := request.URL.Query()
	// Setting max-results very high to try and get all branches
	q.Set("max-results", "10000")
	request.URL.RawQuery = q.Encode()

	branchResponse := BranchesResponse{}
	response, err := pb.client.Do(request, &branchResponse)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Listing VCS branches for %s returned %s", planKey, response.Status)}
	}

	vcsBranches := make([]string, len(branchResponse.Branches.BranchList))

	for i, b := range branchResponse.Branches.BranchList {
		vcsBranches[i] = b.Name
	}

	return vcsBranches, response, err
}
