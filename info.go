package bamboo

import (
	"fmt"
	"net/http"
)

// InfoService retrieves general server information
type InfoService service

// ServerInfo represents the general
// server information exposed by Bamboo
type ServerInfo struct {
	Version     string `json:"version,omitempty"`
	Edition     string `json:"edition,omitempty"`
	BuildDate   string `json:"buildDate,omitempty"`
	BuildNumber string `json:"buildNumber,omitempty"`
	State       string `json:"state,omitempty"`
}

// Info fetches the server information for the Bamboo server
func (i *InfoService) Info() (*ServerInfo, *http.Response, error) {
	u := "info.json"
	request, err := i.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	serverInfo := &ServerInfo{}
	response, err := i.client.Do(request, serverInfo)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Request for server info returned %d", response.StatusCode)}
	}

	return serverInfo, response, nil
}
