package bamboo

import (
	"fmt"
	"net/http"
)

// CloneService handles the cloning of one Bamboo resource to another.
type CloneService service

// ClonePlan clones a plan from one project to another given the full project-plan key for
// the source plan and the destination plan. The destination project does not need to be in
// the same project as the source plan. Returns a Plan struct of the resulting plan.
func (c *CloneService) ClonePlan(srcKey, dstKey string) (*Plan, *http.Response, error) {
	var u string
	if !emptyStrings(srcKey, dstKey) {
		u = fmt.Sprintf("clone/%s:%s.json", srcKey, dstKey)
	} else {
		return nil, nil, &simpleError{"Source key and/or destination key cannot be empty strings"}
	}

	request, err := c.client.NewRequest(http.MethodPut, u, nil)
	if err != nil {
		return nil, nil, err
	}

	clonedPlan := Plan{}
	response, err := c.client.Do(request, &clonedPlan)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Clone returned %d", response.StatusCode)}
	}

	return &clonedPlan, response, nil
}
