package bamboo_test

import (
	"os"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

var (
	// bambooCLI for all other tests
	bambooClient *bamboo.Client
)

func init() {
	bambooClient = bamboo.NewSimpleClient(nil, os.Getenv("BAMBOO_USERNAME"), os.Getenv("BAMBOO_PASSWORD"))
	// May need to set the URL to that of the Docker machine
	//bambooCLI.SetURL("http://192.168.99.100:8085")
}

func TestSetURL(t *testing.T) {

}

func TestNewSimpleClient(t *testing.T) {

}

func TestNewRequest(t *testing.T) {

}

func TestDo(t *testing.T) {

}
