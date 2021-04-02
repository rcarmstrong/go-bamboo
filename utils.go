package bamboo

import (
	"fmt"
	"io"
	"io/ioutil"
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
	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		_, _ = io.CopyN(ioutil.Discard, response.Body, 512)
		_ = response.Body.Close()
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("%s: %s - %q", msg, response.Status, body)
}

// Pagination used to specify the start and limit indexes of a paginated API resource
type Pagination struct {
	Start int
	Limit int
}
