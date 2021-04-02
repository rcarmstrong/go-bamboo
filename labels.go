package bamboo

import (
	"fmt"
	"net/http"
)

// LabelService handles communication with the labels on a plan result
type LabelService service

type ILabelService interface {
	AddLabel(label *Label) (bool, *http.Response, error)
}

// Label is a single label on a result
type Label struct {
	Name      string `json:"name"`
	ResultKey string `json:"-"`
}

func (lb Label) isEmpty() bool {
	return lb.ResultKey == "" || lb.Name == ""
}

// AddLabel will add a label to the given result.
func (c *LabelService) AddLabel(label *Label) (bool, *http.Response, error) {
	if label == nil || label.isEmpty() {
		return false, nil, &simpleError{"Label cannot be nil or empty"}
	}
	u := fmt.Sprintf("result/%s/label.json", label.ResultKey)

	request, err := c.client.NewRequest(http.MethodPost, u, label)
	if err != nil {
		return false, nil, err
	}

	request.Header.Add("Accept", "application/json")

	response, err := c.client.Do(request, nil)
	if err != nil {
		return false, response, err
	}

	if !(response.StatusCode == 204) {
		return false, response, &simpleError{fmt.Sprintf("Adding Label to %s returned %s", label.ResultKey, response.Status)}
	}

	return true, response, nil
}
