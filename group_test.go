package bamboo

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGroupPermissionsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(groupPermissionsListStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		_, resp, err := client.Permissions.GroupPermissionsList(tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func groupPermissionsListStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/groups":           true,
		"plan/TEST/groups":        true,
		"repository/TEST/groups":  true,
		"project/TEST/groups":     true,
		"environment/TEST/groups": true,
		"projectplan/TEST/groups": true,
		"deployment/TEST/groups":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestGroupPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(groupPermissionsStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		_, resp, err := client.Permissions.GroupPermissions("testgroup", tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func groupPermissionsStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/groups?name=testgroup":           true,
		"plan/TEST/groups?name=testgroup":        true,
		"repository/TEST/groups?name=testgroup":  true,
		"project/TEST/groups?name=testgroup":     true,
		"environment/TEST/groups?name=testgroup": true,
		"projectplan/TEST/groups?name=testgroup": true,
		"deployment/TEST/groups?name=testgroup":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestSetGroupPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(setGroupPermissionsStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		resp, err := client.Permissions.SetGroupPermissions("testgroup", []string{}, tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func setGroupPermissionsStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/groups/testgroup":           true,
		"plan/TEST/groups/testgroup":        true,
		"repository/TEST/groups/testgroup":  true,
		"project/TEST/groups/testgroup":     true,
		"environment/TEST/groups/testgroup": true,
		"projectplan/TEST/groups/testgroup": true,
		"deployment/TEST/groups/testgroup":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodPut {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestRemoveGroupPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(removeGroupPermissionsStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		resp, err := client.Permissions.RemoveGroupPermissions("testgroup", []string{}, tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func removeGroupPermissionsStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/groups/testgroup":           true,
		"plan/TEST/groups/testgroup":        true,
		"repository/TEST/groups/testgroup":  true,
		"project/TEST/groups/testgroup":     true,
		"environment/TEST/groups/testgroup": true,
		"projectplan/TEST/groups/testgroup": true,
		"deployment/TEST/groups/testgroup":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodDelete {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestAvailableGroupsPermissionsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(availableGroupsPermissionsListStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		_, resp, err := client.Permissions.AvailableGroupsPermissionsList(tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func availableGroupsPermissionsListStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/available-groups":           true,
		"plan/TEST/available-groups":        true,
		"repository/TEST/available-groups":  true,
		"project/TEST/available-groups":     true,
		"environment/TEST/available-groups": true,
		"projectplan/TEST/available-groups": true,
		"deployment/TEST/available-groups":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
