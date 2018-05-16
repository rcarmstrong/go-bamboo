package bamboo_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

func TestGroupPermissionsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(groupPermissionsListStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, resp, err := client.ProjectPlan.GroupPermissionsList("CORE", nil)
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func groupPermissionsListStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/groups" {
		w.WriteHeader(http.StatusBadRequest)
	} else if r.URL.RawQuery != "start=0&limit=25" {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestGroupPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(groupPermissionsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, resp, err := client.ProjectPlan.GroupPermissions("CORE", "test")
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func groupPermissionsStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/groups" {
		w.WriteHeader(http.StatusBadRequest)
	} else if r.URL.RawQuery != "name=test" {
		w.WriteHeader(http.StatusBadRequest)
	}

	group := bamboo.GroupProjectPlanResponse{
		Results: []bamboo.Group{
			bamboo.Group{
				Name: "test",
			},
		},
	}

	data, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}

	w.Write(data)
}

func TestSetGroupPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(setGroupPermissionsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	permissions := []string{
		bamboo.ReadPermission,
		bamboo.BuildPermission,
		bamboo.WritePermission,
	}

	resp, err := client.ProjectPlan.SetGroupPermissions("CORE", "test", permissions)
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func setGroupPermissionsStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/groups/test" {
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

func TestRemoveGroupPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(removeGroupPermissionsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	permissions := []string{
		bamboo.ReadPermission,
		bamboo.BuildPermission,
		bamboo.WritePermission,
	}

	resp, err := client.ProjectPlan.RemoveGroupPermissions("CORE", "test", permissions)
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func removeGroupPermissionsStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/groups/test" {
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

func TestAvailableGroupPermissionsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(availableGroupPermissionsListStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, resp, err := client.ProjectPlan.AvailableGroupPermissionsList("CORE", nil)
	if err != nil {
		log.Println(resp.Status)
		t.Error(err)
	}
}

func availableGroupPermissionsListStub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/available-groups" {
		w.WriteHeader(http.StatusBadRequest)
	} else if r.URL.RawQuery != "start=0&limit=25" {
		w.WriteHeader(http.StatusBadRequest)
	}
}
