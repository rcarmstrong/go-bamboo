package bamboo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var serverState = &TransitionStateInfo{
	ServerInfo: ServerInfo{
		State:             PausedState,
		ReindexInProgress: false,
	},
	SetByUser: "test",
}

func TestStateTransitions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(transitionServerStateStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	var testCases = []struct {
		expectedState string
		function      func() (*TransitionStateInfo, *http.Response, error)
	}{
		{PausedState, client.Server.Pause},
		{RunningState, client.Server.Resume},
		{PreparingForRestartState, client.Server.PrepareForRestart},
		{ReadyForRestartState, client.Server.Resume},
	}

	for _, c := range testCases {
		transitionStateInfo, _, err := c.function()
		if err != nil {
			t.Error(err)
		}

		if transitionStateInfo.State != c.expectedState {
			t.Error(fmt.Sprintf("Server state %s does not equal expected state of %s", transitionStateInfo.State, c.expectedState))
		}
	}
}

func TestReindex(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(reindexServerStateStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	var testCases = []struct {
		expected bool
		function func() (*ReindexState, *http.Response, error)
	}{
		{true, client.Server.Reindex},
		{true, client.Server.ReindexStatus},
	}

	for _, c := range testCases {
		reindexState, _, err := c.function()
		if err != nil {
			t.Error(err)
		}

		if reindexState.ReindexInProgress != c.expected {
			t.Error(fmt.Sprintf("Reindex method returned %t when %t was expected", reindexState.ReindexInProgress, c.expected))
		}
	}
}

func transitionServerStateStub(w http.ResponseWriter, r *http.Request) {
	method := strings.Split(strings.Split(r.URL.String(), ".")[0], "/")[5]

	switch method {
	case "pause":
		serverState.State = PausedState
	case "resume":
		if serverState.State == PausedState {
			serverState.State = RunningState
		} else {
			serverState.State = ReadyForRestartState
		}
	case "prepareForRestart":
		serverState.State = PreparingForRestartState
	default:
		serverState.State = "Unknown"
	}

	bytes, err := json.Marshal(serverState)
	if err != nil {
		panic(err)
	}

	w.Write(bytes)
}

func reindexServerStateStub(w http.ResponseWriter, r *http.Request) {
	resp := ReindexState{
		ReindexInProgress: true,
		ReindexPending:    true,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(202)
	case "GET":
		w.WriteHeader(200)
	}

	w.Write(bytes)
}
