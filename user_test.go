package bamboo_test

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	bamboo "github.com/rcarmstrong/go-bamboo"
// )

// func TestUserPermissionsList(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(userPermissionsListStub))
// 	defer ts.Close()

// 	client := bamboo.NewSimpleClient(nil, "", "")
// 	client.SetURL(ts.URL)

// 	_, resp, err := client.Permissions.ProjectPlan.UserPermissionsList("CORE")
// 	if err != nil {
// 		log.Println(resp.Status)
// 		t.Error(err)
// 	}
// }

// func userPermissionsListStub(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/users" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }

// func TestUserPermissions(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(userPermissionsStub))
// 	defer ts.Close()

// 	client := bamboo.NewSimpleClient(nil, "", "")
// 	client.SetURL(ts.URL)

// 	_, resp, err := client.Permissions.ProjectPlan.UserPermissions("CORE", "test")
// 	if err != nil {
// 		log.Println(resp.Status)
// 		t.Error(err)
// 	}
// }

// func userPermissionsStub(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/users" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	} else if r.URL.RawQuery != "name=test" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	user := bamboo.UserPermissionsResponse{
// 		Results: []bamboo.User{
// 			bamboo.User{
// 				Name: "test",
// 			},
// 		},
// 	}

// 	userData, err := json.Marshal(user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.Write(userData)
// }

// func TestSetUserPermissions(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(setUserPermissionsStub))
// 	defer ts.Close()

// 	client := bamboo.NewSimpleClient(nil, "", "")
// 	client.SetURL(ts.URL)

// 	permissions := []string{
// 		bamboo.ReadPermission,
// 		bamboo.BuildPermission,
// 		bamboo.WritePermission,
// 	}

// 	resp, err := client.Permissions.ProjectPlan.SetUserPermissions("CORE", "test", permissions)
// 	if err != nil {
// 		log.Println(resp.Status)
// 		t.Error(err)
// 	}
// }

// func setUserPermissionsStub(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/users/test" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	} else if r.Method != "PUT" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	permissions := []string{}
// 	bytes, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	err = json.Unmarshal(bytes, &permissions)
// 	if err != nil {
// 		panic(err)
// 	}
// 	status := http.StatusBadRequest
// 	for _, p := range permissions {
// 		switch p {
// 		case bamboo.ReadPermission:
// 			status = http.StatusNoContent
// 		case bamboo.WritePermission:
// 			status = http.StatusNoContent
// 		case bamboo.BuildPermission:
// 			status = http.StatusNoContent
// 		default:
// 			status = http.StatusBadRequest
// 		}
// 	}

// 	w.WriteHeader(status)
// }

// func TestRemoveUserPermissions(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(removeUserPermissionsStub))
// 	defer ts.Close()

// 	client := bamboo.NewSimpleClient(nil, "", "")
// 	client.SetURL(ts.URL)

// 	permissions := []string{
// 		bamboo.ReadPermission,
// 		bamboo.BuildPermission,
// 		bamboo.WritePermission,
// 	}

// 	resp, err := client.Permissions.ProjectPlan.RemoveUserPermissions("CORE", "test", permissions)
// 	if err != nil {
// 		log.Println(resp.Status)
// 		t.Error(err)
// 	}
// }

// func removeUserPermissionsStub(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/users/test" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	} else if r.Method != "DELETE" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	permissions := []string{}
// 	bytes, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	err = json.Unmarshal(bytes, &permissions)
// 	if err != nil {
// 		panic(err)
// 	}
// 	status := http.StatusBadRequest
// 	for _, p := range permissions {
// 		switch p {
// 		case bamboo.ReadPermission:
// 			status = http.StatusNoContent
// 		case bamboo.WritePermission:
// 			status = http.StatusNoContent
// 		case bamboo.BuildPermission:
// 			status = http.StatusNoContent
// 		default:
// 			status = http.StatusBadRequest
// 		}
// 	}

// 	w.WriteHeader(status)
// }

// func TestAvailableUserPermissionsList(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(availableUserPermissionsListStub))
// 	defer ts.Close()

// 	client := bamboo.NewSimpleClient(nil, "", "")
// 	client.SetURL(ts.URL)

// 	_, resp, err := client.Permissions.ProjectPlan.AvailableUserPermissionsList("CORE")
// 	if err != nil {
// 		log.Println(resp.Status)
// 		t.Error(err)
// 	}
// }

// func availableUserPermissionsListStub(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/rest/api/latest/permissions/projectplan/CORE/available-users" {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }
