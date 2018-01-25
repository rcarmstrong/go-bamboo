package bamboo

import (
	"fmt"
	"net/http"
)

// CommentService handles communication with the comments on a plan result
type CommentService service

// Comment is a single comment on a result
type Comment struct {
	Content   string `json:"content"`
	ResultKey string `json:"-"`
}

func (cm Comment) isEmpty() bool {
	return cm.ResultKey == "" || cm.Content == ""
}

// AddComment will add a comment to the given result.
func (c *CommentService) AddComment(comment *Comment) (bool, *http.Response, error) {
	if comment == nil || comment.isEmpty() {
		return false, nil, &simpleError{"Comment cannot be nil or empty"}
	}
	u := fmt.Sprintf("result/%s/comment.json", comment.ResultKey)

	request, err := c.client.NewRequest("POST", u, comment)
	if err != nil {
		return false, nil, err
	}

	request.Header.Add("Accept", "application/json")

	response, err := c.client.Do(request, nil)
	if err != nil {
		return false, response, err
	}

	if !(response.StatusCode == 204) {
		return false, response, &simpleError{fmt.Sprintf("Adding comment to %s returned %s", comment.ResultKey, response.Status)}
	}

	return true, response, nil
}
