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
	testClient := bamboo.NewSimpleClient(nil, "myusername", "mypassword")
	cases := []struct {
		testURL     string
		expected    string
		shouldError bool
	}{
		{
			testURL:     "http://localhost:8085",
			expected:    "http://localhost:8085/rest/api/latest/",
			shouldError: false,
		},
		{
			testURL:     "fuzzybunnyslippers",
			expected:    "fuzzybunnyslippers/rest/api/latest",
			shouldError: true,
		},
	}

	for _, c := range cases {
		err := testClient.SetURL(c.testURL)
		if err != nil {
			if c.shouldError {
				t.Logf("%s expected to throw an err %s", c.testURL, err.Error())
			} else {
				t.Error(err)
			}
		} else {
			if testClient.BaseURL.String() != c.expected {
				t.Errorf("Client BaseURL is %s when we expected %s", testClient.BaseURL.String(), c.expected)
			}
		}
	}
}

func TestNewSimpleClient(t *testing.T) {

}

func TestNewRequest(t *testing.T) {

}

func TestDo(t *testing.T) {

}
