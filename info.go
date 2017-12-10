package bamboo

import "fmt"

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
func (i *InfoService) Info() (*ServerInfo, error) {
	u := "info.json"
	req, err := i.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	serverInfo := &ServerInfo{}
	resp, err := i.client.Do(req, serverInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == 200) {
		return nil, &simpleError{fmt.Sprintf("Request for server info returned %d", resp.StatusCode)}
	}

	return serverInfo, nil
}
