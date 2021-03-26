package test

import (
	"github.com/lotos2512/bamboo"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserPermissionsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(userPermissionsListStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		_, resp, err := client.Permissions.UserPermissionsList(tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func userPermissionsListStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/users":           true,
		"plan/TEST/users":        true,
		"repository/TEST/users":  true,
		"project/TEST/users":     true,
		"environment/TEST/users": true,
		"projectplan/TEST/users": true,
		"deployment/TEST/users":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestUserPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(userPermissionsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		_, resp, err := client.Permissions.UserPermissions("testuser", tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func userPermissionsStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/users?name=testuser":           true,
		"plan/TEST/users?name=testuser":        true,
		"repository/TEST/users?name=testuser":  true,
		"project/TEST/users?name=testuser":     true,
		"environment/TEST/users?name=testuser": true,
		"projectplan/TEST/users?name=testuser": true,
		"deployment/TEST/users?name=testuser":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestSetUserPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(setUserPermissionsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		resp, err := client.Permissions.SetUserPermissions("testuser", []string{}, tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func setUserPermissionsStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/users/testuser":           true,
		"plan/TEST/users/testuser":        true,
		"repository/TEST/users/testuser":  true,
		"project/TEST/users/testuser":     true,
		"environment/TEST/users/testuser": true,
		"projectplan/TEST/users/testuser": true,
		"deployment/TEST/users/testuser":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodPut {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestRemoveUserPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(removeUserPermissionsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		resp, err := client.Permissions.RemoveUserPermissions("testuser", []string{}, tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func removeUserPermissionsStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/users/testuser":           true,
		"plan/TEST/users/testuser":        true,
		"repository/TEST/users/testuser":  true,
		"project/TEST/users/testuser":     true,
		"environment/TEST/users/testuser": true,
		"projectplan/TEST/users/testuser": true,
		"deployment/TEST/users/testuser":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodDelete {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestAvailableUserPermissionsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(availableUserPermissionsListStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	for _, tc := range permissionsTestCases {
		_, resp, err := client.Permissions.AvailableUsersPermissionsList(tc)
		if err != nil {
			log.Println(resp.Status)
			t.Error(err)
		}
	}
}

func availableUserPermissionsListStub(w http.ResponseWriter, r *http.Request) {
	var expectedURLs = map[string]bool{
		"global/available-users":           true,
		"plan/TEST/available-users":        true,
		"repository/TEST/available-users":  true,
		"project/TEST/available-users":     true,
		"environment/TEST/available-users": true,
		"projectplan/TEST/available-users": true,
		"deployment/TEST/available-users":  true,
	}

	check := strings.Split(r.URL.String(), "permissions/")[1]
	if expectedURLs[check] && r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
