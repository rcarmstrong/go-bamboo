package bamboo

import (
	"fmt"
	"net/http"
)

// InfoService retrieves server information
type InfoService service

// BuildInfo represents the build information of the Bamboo server
type BuildInfo struct {
	Version     string `json:"version,omitempty"`
	Edition     string `json:"edition,omitempty"`
	BuildDate   string `json:"buildDate,omitempty"`
	BuildNumber string `json:"buildNumber,omitempty"`
	State       string `json:"state,omitempty"`
}

// ServerInfo contains information on the Bamboo server
type ServerInfo struct {
	State             string `json:"state"`
	ReindexInProgress bool   `json:"reindexInProgress"`
}

func (s *ServerInfo) isRunning() bool {
	return s.State == "RUNNING"
}

// BuildInfo fetches the build information of the Bamboo server
func (i *InfoService) BuildInfo() (*BuildInfo, *http.Response, error) {
	u := "info.json"
	request, err := i.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	buildInfo := &BuildInfo{}
	response, err := i.client.Do(request, buildInfo)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Request for server build info returned %d", response.StatusCode)}
	}

	return buildInfo, response, nil
}

// ServerInfo fetches the Bamboo server information
func (i *InfoService) ServerInfo() (*ServerInfo, *http.Response, error) {
	u := "server.json"
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
