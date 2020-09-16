package bamboo

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRolePermissionsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rolePermissionsListStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		_, resp, err := client.Permissions.RolePermissionsList(tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func rolePermissionsListStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/roles":           true,
		"plan/TEST/roles":        true,
		"repository/TEST/roles":  true,
		"project/TEST/roles":     true,
		"environment/TEST/roles": true,
		"projectplan/TEST/roles": true,
		"deployment/TEST/roles":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestSetLoggedInUserPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(setLoggedInUserPermissionsStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		resp, err := client.Permissions.SetLoggedInUsersPermissions([]string{}, tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func setLoggedInUserPermissionsStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/roles/LOGGED_IN":           true,
		"plan/TEST/roles/LOGGED_IN":        true,
		"repository/TEST/roles/LOGGED_IN":  true,
		"project/TEST/roles/LOGGED_IN":     true,
		"environment/TEST/roles/LOGGED_IN": true,
		"projectplan/TEST/roles/LOGGED_IN": true,
		"deployment/TEST/roles/LOGGED_IN":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodPut {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestRemoveLoggedInUsersPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(removeLoggedInUsersPermissionsStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		resp, err := client.Permissions.RemoveLoggedInUsersPermissions([]string{}, tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func removeLoggedInUsersPermissionsStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/roles/LOGGED_IN":           true,
		"plan/TEST/roles/LOGGED_IN":        true,
		"repository/TEST/roles/LOGGED_IN":  true,
		"project/TEST/roles/LOGGED_IN":     true,
		"environment/TEST/roles/LOGGED_IN": true,
		"projectplan/TEST/roles/LOGGED_IN": true,
		"deployment/TEST/roles/LOGGED_IN":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodDelete {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestSetAnonymousReadPermission(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(setAnonymousReadPermissionStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		resp, err := client.Permissions.SetAnonymousReadPermission(tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func setAnonymousReadPermissionStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/roles/ANONYMOUS":           true,
		"plan/TEST/roles/ANONYMOUS":        true,
		"repository/TEST/roles/ANONYMOUS":  true,
		"project/TEST/roles/ANONYMOUS":     true,
		"environment/TEST/roles/ANONYMOUS": true,
		"projectplan/TEST/roles/ANONYMOUS": true,
		"deployment/TEST/roles/ANONYMOUS":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodPut {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestRemoveAnonymousReadPermission(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(removeAnonymousReadPermissionStub))
	defer ts.Close()

	client := NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		resp, err := client.Permissions.RemoveAnonymousReadPermission(tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func removeAnonymousReadPermissionStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/roles/ANONYMOUS":           true,
		"plan/TEST/roles/ANONYMOUS":        true,
		"repository/TEST/roles/ANONYMOUS":  true,
		"project/TEST/roles/ANONYMOUS":     true,
		"environment/TEST/roles/ANONYMOUS": true,
		"projectplan/TEST/roles/ANONYMOUS": true,
		"deployment/TEST/roles/ANONYMOUS":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodDelete {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
