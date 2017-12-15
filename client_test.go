package bamboo_test

import (
	"net/http"
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
	testClient := bamboo.NewSimpleClient(nil, "myusername", "mypassword")
	compareClient := struct {
		client      *http.Client
		baseURL     string
		simpleCreds *bamboo.SimpleCredentials
	}{
		client:  http.DefaultClient,
		baseURL: "http://localhost:8085/rest/api/latest/",
		simpleCreds: &bamboo.SimpleCredentials{
			Username: "myusername",
			Password: "mypassword",
		},
	}

	if compareClient.baseURL != testClient.BaseURL.String() {
		t.Errorf("Expected client BaseURL to be %s but got %s", compareClient.baseURL, testClient.BaseURL.String())
	}

	if compareClient.simpleCreds.Username != testClient.SimpleCreds.Username {
		t.Errorf("Expected client Username to be %s but got %s", compareClient.simpleCreds.Username, testClient.SimpleCreds.Username)
	}

	if compareClient.simpleCreds.Password != testClient.SimpleCreds.Password {
		t.Errorf("Expected client Username to be %s but got %s", compareClient.simpleCreds.Password, testClient.SimpleCreds.Password)
	}

	//TODO compair http.Clients
}

func TestNewRequest(t *testing.T) {

}

func TestDo(t *testing.T) {

}
