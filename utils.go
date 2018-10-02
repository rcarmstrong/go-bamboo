package bamboo

import (
	"fmt"
	"net/http"
)

func emptyStrings(strings ...string) bool {
	for _, s := range strings {
		if s == "" {
			return true
		}
	}
	return false
}

func newRespErr(response *http.Response, msg string) error {
	body := []byte{}
	response.Body.Read(body)
	return fmt.Errorf("%s: %s - %q", msg, response.Status, body)
}

// Pagination used to specify the start and limit indexes of a paginated API resource
type Pagination struct {
	Start int
	Limit int
}
