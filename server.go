package bamboo

import (
	"fmt"
	"net/http"
)

// RunningState is the state of a running Bamboo server
const RunningState string = "RUNNING"

// PausingState is when the Bamboo server is in the process of being paused
const PausingState string = "PAUSING"

// PausedState is the state of a paused Bamboo server
const PausedState string = "PAUSED"

// ReadyForRestartState is the state of a Bamboo server ready to be restarted
const ReadyForRestartState string = "READY_FOR_RESTART"

// PreparingForRestartState is the state of a Bamboo server preparing to be restarted
const PreparingForRestartState string = "PREPARING_FOR_RESTART"

// ServerService exposes server operations
type ServerService service

// TransitionStateInfo represents the server state response after a server operation is preformed.
type TransitionStateInfo struct {
	ServerInfo
	SetByUser string `json:"setByUser"`
}

// ReindexState represents the state of a server reindex.
// ReindexInProgress - true if a reindex is in progress otherwise false
// ReindexPending - reindex is required (i.e. it failed before or some upgrade task asked for it)
type ReindexState struct {
	ReindexInProgress bool `json:"reindexInProgress"`
	ReindexPending    bool `json:"reindexPending"`
}

// Pause will move the Bamboo server to the PAUSED state.
// The PAUSED state only prevents new builds from being scheduled. Change detection and
// other server operations will continue to run.
func (s *ServerService) Pause() (*TransitionStateInfo, *http.Response, error) {
	u := "server/pause.json"
	request, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	state := &TransitionStateInfo{}
	response, err := s.client.Do(request, state)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Server pause returned %d", response.StatusCode)}
	}

	return state, response, nil
}

// Resume will move the Bamboo server to either the RUNNING or READY_FOR_RESTART state.
// The RUNNING state means the server was PAUSED and builds will resume.
// The READY_FOR_RESTART state means exactly what the name suggests and builds will not resume
// until the server is restarted.
func (s *ServerService) Resume() (*TransitionStateInfo, *http.Response, error) {
	u := "server/resume.json"
	request, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	state := &TransitionStateInfo{}
	response, err := s.client.Do(request, state)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Server resume returned %d", response.StatusCode)}
	}

	return state, response, nil
}

// PrepareForRestart will move the Bamboo server to the PREPARING_FOR_RESTART state.
// Change detection, indexing, ec2 instance ordering etc. are stopped to allow for a server restart.
func (s *ServerService) PrepareForRestart() (*TransitionStateInfo, *http.Response, error) {
	u := "server/prepareForRestart.json"
	request, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, nil, err
	}

	state := &TransitionStateInfo{}
	response, err := s.client.Do(request, state)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Server prepare for restart returned %d", response.StatusCode)}
	}

	return state, response, nil
}

// Reindex will start a server reindex
func (s *ServerService) Reindex() (*ReindexState, *http.Response, error) {
	u := "reindex"
	request, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	state := &ReindexState{}
	response, err := s.client.Do(request, state)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 202) {
		return nil, response, &simpleError{fmt.Sprintf("Server reindex returned %d", response.StatusCode)}
	}

	return state, response, nil
}

// ReindexStatus will start a server reindex
func (s *ServerService) ReindexStatus() (*ReindexState, *http.Response, error) {
	u := "reindex"
	request, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	state := &ReindexState{}
	response, err := s.client.Do(request, state)
	if err != nil {
		return nil, response, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("Request for reindex status returned %d", response.StatusCode)}
	}

	return state, response, nil
}
