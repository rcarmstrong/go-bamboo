package bamboo

import (
	"fmt"
	"io/ioutil"
)

// ProjectService handles communication with the project related methods
type ProjectService service

// ListProjectsOptions allows you to pass a project name to get information on
type ListProjectsOptions struct {
	Project string
}

// ProjectResponse wraps all the information about the projects
type ProjectResponse struct {
	List Projects `json:"projects"`
}

// Projects is a collection of project elements
type Projects struct {
	List []Project `json:"project"`
}

// Project is a single project definition
type Project struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// ListProjects lists all projects unless ListProjectsOptions are passed
func (p *ProjectService) ListProjects(opt *ListProjectsOptions) ([]Project, error) {
	u := "project"
	if opt != nil {
		if !emptyStrings(opt.Project) {
			u += fmt.Sprintf("/%s.json", opt.Project)
		}
	} else {
		u += ".json"
	}

	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(req.URL.String())

	projectResp := ProjectResponse{}
	resp, err := p.client.Do(req, &projectResp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

	if !(resp.StatusCode == 200) {
		return nil, &simpleError{fmt.Sprintf("List projects returned %s", resp.Status)}
	}

	fmt.Println(len(projectResp.List.List))

	return projectResp.List.List, nil
}
