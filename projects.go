package bamboo

import (
	"fmt"
	"net/http"
)

// ProjectService handles communication with the project related methods
type ProjectService service

// ProjectResponse the REST response from the server
type ProjectResponse struct {
	*ServiceMetadata
	Projects *Projects `json:"projects"`
}

// Projects is a collection of project elements
type Projects struct {
	*CollectionMetadata
	ProjectList []*Project `json:"project"`
}

// Project is a single project definition
type Project struct {
	Key         string       `json:"key,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Link        *ServiceLink `json:"link,omitempty"`
}

// ProjectInformation is the information for a single project
type ProjectInformation struct {
	Key         string                   `json:"key,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Description string                   `json:"description,omitempty"`
	NumPlans    *ProjectPlansInformation `json:"plans"`
}

// ProjectPlansInformation holds the number of plans in a project
type ProjectPlansInformation struct {
	Size int `json:"size,omitempty"`
}

// ProjectInfo get the information on the specific project
func (p *ProjectService) ProjectInfo(projectKey string) (*ProjectInformation, *http.Response, error) {
	var u string
	if !emptyStrings(projectKey) {
		u = fmt.Sprintf("project/%s.json", projectKey)
	} else {
		return nil, nil, &simpleError{fmt.Sprintf("Project key cannot be an empty string")}
	}

	request, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	projectInfo := ProjectInformation{}
	response, err := p.client.Do(request, &projectInfo)
	if err != nil {
		return nil, nil, err
	}

	return &projectInfo, response, nil
}

// ListProjects lists all projects
func (p *ProjectService) ListProjects() ([]*Project, *http.Response, error) {
	u := "project.json"

	request, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	projectResp := ProjectResponse{}
	response, err := p.client.Do(request, &projectResp)
	if err != nil {
		return nil, nil, err
	}

	if !(response.StatusCode == 200) {
		return nil, response, &simpleError{fmt.Sprintf("List projects returned %s", response.Status)}
	}

	return projectResp.Projects.ProjectList, response, nil
}
