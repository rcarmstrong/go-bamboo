package bamboo_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

var serverState = &bamboo.TransitionStateInfo{
	bamboo.ServerInfo{
		State:             bamboo.PausedState,
		ReindexInProgress: false,
	},
	"test",
}

func TestStateTransitions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(transitionServerStateStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	var testCases = []struct {
		expectedState string
		function      func() (*bamboo.TransitionStateInfo, *http.Response, error)
	}{
		{bamboo.PausedState, client.Server.Pause},
		{bamboo.RunningState, client.Server.Resume},
		{bamboo.PreparingForRestartState, client.Server.PrepareForRestart},
		{bamboo.ReadyForRestartState, client.Server.Resume},
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

func transitionServerStateStub(w http.ResponseWriter, r *http.Request) {
	method := strings.Split(strings.Split(r.URL.String(), ".")[0], "/")[5]

	switch method {
	case "pause":
		serverState.State = bamboo.PausedState
	case "resume":
		if serverState.State == bamboo.PausedState {
			serverState.State = bamboo.RunningState
		} else {
			serverState.State = bamboo.ReadyForRestartState
		}
	case "prepareForRestart":
		serverState.State = bamboo.PreparingForRestartState
	default:
		serverState.State = "Unknown"
	}

	bytes, err := json.Marshal(serverState)
	if err != nil {
		panic(err)
	}

	w.Write(bytes)
}
