package bamboo_test

import (
	"testing"
)

func TestInfo(t *testing.T) {
	info, resp, err := bambooClient.Server.Info()
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Getting Server Info returned %s", resp.Status)
	}

	if info.Version == "" && info.BuildNumber == "" && info.State == "" {
		t.Errorf("Server responded with vital fields empty: %+v", info)
	}
}
