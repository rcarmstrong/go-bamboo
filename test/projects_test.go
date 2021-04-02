package test

import (
	"encoding/json"
	"github.com/lotos2512/bamboo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(unauthorizedStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, response, err := client.Projects.ProjectInfo("ABC")
	assert.Error(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestProjectInfo2(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(projectInfoStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	result, response, err := client.Projects.ProjectInfo("ABC")
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.NotNil(t, result)
	assert.Equal(t, 0, result.NumPlans.Size)
}

func projectInfoStub(w http.ResponseWriter, r *http.Request) {
	resp := bamboo.ProjectInformation{Key: "ABC", Name: "abc",
		Description: "complex",
		NumPlans:    &bamboo.ProjectPlansInformation{Size: 0}}

	bytes, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func TestProjectPlans(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(unauthorizedStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, response, err := client.Projects.ProjectPlans("ABC")
	assert.Error(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestProjectPlans2(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(projectPlansStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	result, response, err := client.Projects.ProjectPlans("ABC")
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
}

func projectPlansStub(w http.ResponseWriter, r *http.Request) {
	resp := bamboo.PlanResponse{Plans: &bamboo.Plans{PlanList: []*bamboo.Plan{&bamboo.Plan{"test project", "TPRJ", "", true, nil,
		"ABC-TPRJ", "test project", &bamboo.PlanKey{"ABC-TPRJ"}}}}}
	bytes, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func TestListProjects(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(unauthorizedStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	_, response, err := client.Projects.ListProjects()
	assert.Error(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestListProjects2(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(listProjectsStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	result, response, err := client.Projects.ListProjects()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
}

func listProjectsStub(w http.ResponseWriter, r *http.Request) {
	resp := bamboo.ProjectResponse{Projects: &bamboo.Projects{ProjectList: []*bamboo.Project{&bamboo.Project{"ABC", "first project", "long description", nil}}}}

	bytes, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func unauthorizedStub(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}
