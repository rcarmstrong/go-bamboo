package bamboo

import (
	"fmt"
	"net/http"
)

// CommentService handles communication with the comments on a plan result
type CommentService service

type Comment struct {
	Content string
}

// AddComment will add a comment to the given result.
func (c *CommentService) AddComment(resultKey, comment string) (bool, *http.Response, error) {
	if resultKey == "" {
		return true, nil, &simpleError{"resultKey cannot be blank"}
	}
	u := fmt.Sprintf("result/%s/comment", resultKey)
	payload := Comment{comment}

	request, err := c.client.NewRequest("POST", u, payload)
	if err != nil {
		return false, nil, err
	}

	response, err := c.client.Do(request, nil)
	if err != nil {
		return false, response, err
	}

	if !(response.StatusCode == 200) {
		return false, response, &simpleError{fmt.Sprintf("Adding comment to %s returned %d", resultKey, response.StatusCode)}
	}

	return true, response, nil
}
