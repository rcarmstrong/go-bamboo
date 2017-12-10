package bamboo_test

import (
	"os"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

var (
	bambooCLI *bamboo.Client
)

func init() {
	bambooCLI = bamboo.NewSimpleClient(nil, os.Getenv("BAMBOO_USERNAME"), os.Getenv("BAMBOO_PASSWORD"))
}

func TestInfo(t *testing.T) {
	info, err := bambooCLI.Server.Info()
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", info)
}
