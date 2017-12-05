package bamboo

import (
	"fmt"
)

// ProjectService handles communication with the project related methods
type ProjectService service

// ProjectResponse the REST response from the server
type ProjectResponse struct {
	Response Projects `json:"projects"`
}

// Projects is a collection of project elements
type Projects struct {
	ProjectList []Project `json:"project"`
}

// Project is a single project definition
type Project struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// ProjectInformation is the information for a single project
type ProjectInformation struct {
	Key         string                  `json:"key,omitempty"`
	Name        string                  `json:"name,omitempty"`
	Description string                  `json:"description,omitempty"`
	NumPlans    ProjectPlansInformation `json:"plans"`
}

// ProjectPlansInformation holds the number of plans in a project
type ProjectPlansInformation struct {
	Size int `json:"size,omitempty"`
}

// ProjectInfo get the information on the specific project
func (p *ProjectService) ProjectInfo(projectKey string) (*ProjectInformation, error) {
	var u string
	if !emptyStrings(projectKey) {
		u = fmt.Sprintf("project/%s.json", projectKey)
	} else {
		return nil, &simpleError{fmt.Sprintf("Project key cannot be an empty string")}
	}

	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	projectInfo := ProjectInformation{}
	resp, err := p.client.Do(req, &projectInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &projectInfo, nil
}

// ListProjects lists all projects
func (p *ProjectService) ListProjects() ([]Project, error) {
	u := "project.json"

	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	projectResp := ProjectResponse{}
	resp, err := p.client.Do(req, &projectResp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == 200) {
		return nil, &simpleError{fmt.Sprintf("List projects returned %s", resp.Status)}
	}

	return projectResp.Response.ProjectList, nil
}
