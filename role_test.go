package bamboo_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

func TestRolePermissionsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rolePermissionsListStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, resp, err := client.ProjectPlan.RolePermissionsList("CORE")
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func rolePermissionsListStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/roles" {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestSetLoggedInUserPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(setLoggedInUserPermissionsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	permissions := []string{
		bamboo.ReadPermission,
		bamboo.BuildPermission,
		bamboo.WritePermission,
	}

	resp, err := client.ProjectPlan.SetLoggedInUserPermissions("CORE", permissions)
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func setLoggedInUserPermissionsStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != fmt.Sprintf("/rest/api/latest/permissions/projectplan/CORE/roles/%s", bamboo.LoggedInRole) {
		w.WriteHeader(http.StatusBadRequest)
	} else if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
	}

	permissions := []string{}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(bytes, &permissions)
	if err != nil {
		panic(err)
	}
	status := http.StatusBadRequest
	for _, p := range permissions {
		switch p {
		case bamboo.ReadPermission:
			status = http.StatusNoContent
		case bamboo.WritePermission:
			status = http.StatusNoContent
		case bamboo.BuildPermission:
			status = http.StatusNoContent
		default:
			status = http.StatusBadRequest
		}
	}

	w.WriteHeader(status)
}

func TestRemoveLoggedInUsersPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(removeLoggedInUsersPermissionsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	permissions := []string{
		bamboo.ReadPermission,
		bamboo.BuildPermission,
		bamboo.WritePermission,
	}

	resp, err := client.ProjectPlan.RemoveLoggedInUsersPermissions("CORE", permissions)
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func removeLoggedInUsersPermissionsStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != fmt.Sprintf("/rest/api/latest/permissions/projectplan/CORE/roles/%s", bamboo.LoggedInRole) {
		w.WriteHeader(http.StatusBadRequest)
	} else if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
	}

	permissions := []string{}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(bytes, &permissions)
	if err != nil {
		panic(err)
	}
	status := http.StatusBadRequest
	for _, p := range permissions {
		switch p {
		case bamboo.ReadPermission:
			status = http.StatusNoContent
		case bamboo.WritePermission:
			status = http.StatusNoContent
		case bamboo.BuildPermission:
			status = http.StatusNoContent
		default:
			status = http.StatusBadRequest
		}
	}

	w.WriteHeader(status)
}

func TestSetAnonymousReadPermission(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(setAnonymousReadPermissionStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	resp, err := client.ProjectPlan.SetAnonymousReadPermission("CORE")
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func setAnonymousReadPermissionStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != fmt.Sprintf("/rest/api/latest/permissions/projectplan/CORE/roles/%s", bamboo.AnonymousRole) {
		w.WriteHeader(http.StatusBadRequest)
	} else if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
	}

	permissions := []string{}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(bytes, &permissions)
	if err != nil {
		panic(err)
	}

	if permissions[0] != bamboo.ReadPermission {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusNoContent)
}

func TestRemoveAnonymousReadPermission(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(removeAnonymousReadPermissionStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	resp, err := client.ProjectPlan.RemoveAnonymousReadPermission("CORE")
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func removeAnonymousReadPermissionStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != fmt.Sprintf("/rest/api/latest/permissions/projectplan/CORE/roles/%s", bamboo.AnonymousRole) {
		w.WriteHeader(http.StatusBadRequest)
	} else if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
	}

	permissions := []string{}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(bytes, &permissions)
	if err != nil {
		panic(err)
	}

	if permissions[0] != bamboo.ReadPermission {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusNoContent)
}
