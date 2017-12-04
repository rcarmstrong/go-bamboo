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

	if !(resp.StatusCode == 200) {
		return resp, &simpleError{fmt.Sprintf("Create returned %d", resp.StatusCode)}
	}

	return resp, nil
}
