package bamboo

import (
	"errors"
	"fmt"
	"net/http"
)

type RepositoryService service

type IRepositoryService interface {
	ListRepository(projectKey string, params ListRepositoryParams) (repos []Repository, err error)
	RepositoryScanStatus(params ScanStatusParams) (statusResponse ScanStatusResponse, err error)
}

type Repository struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	RssEnabled bool   `json:"rssEnabled"`
}

type ListRepositoryParams struct {
	RepositoryName string
}

type ScanStatusParams struct {
	ID     uint
	Branch string
}

type ScanStatusResponse struct {
	InProgress bool       `json:"inProgress"`
	SpecsLogs  []SpecsLog `json:"specsLogs"`
}

type SpecImportState string

var (
	SpecImportStateSuccess    SpecImportState = "SUCCESS"
	SpecImportStateError      SpecImportState = "ERROR"
	SpecImportStateInProgress SpecImportState = ""
)

type SpecsLog struct {
	VcsLocationId         uint            `json:"vcsLocationId"`
	Revision              string          `json:"revision"`
	SpecsExecutionDate    uint            `json:"specsExecutionDate"`
	LogFilename           string          `json:"logFilename"`
	SpecImportState       SpecImportState `json:"specImportState"`
	RelativeExecutionDate string          `json:"relativeExecutionDate"`
}

func (p *RepositoryService) ListRepository(projectKey string, params ListRepositoryParams) (repos []Repository, err error) {
	var u string
	if !emptyStrings(projectKey) {
		u = fmt.Sprintf("project/%s/repository", projectKey)
	} else {
		return nil, &simpleError{"Project key cannot be an empty string"}
	}

	request, err := p.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	response, err := p.client.Do(request, &repos)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		return nil,
			errors.New("Getting Project Plans returned: " + response.Status)
	}

	// filter
	if params.RepositoryName != "" {
		for i := range repos {
			if repos[i].Name == params.RepositoryName {
				repos = []Repository{repos[i]}
				break
			}
		}
	}

	return repos, nil
}

func (p *RepositoryService) RepositoryScanStatus(params ScanStatusParams) (statusResponse ScanStatusResponse, err error) {
	var u string
	if params.ID != 0 {
		u = fmt.Sprintf("repository/%d/scan/status", params.ID)
	} else {
		return statusResponse, &simpleError{"ID must be set"}
	}

	request, err := p.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	values := request.URL.Query()
	values.Set("branch", params.Branch)
	values.Set("max-result", "1000")
	request.URL.RawQuery = values.Encode()

	response, err := p.client.Do(request, &statusResponse)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		return statusResponse, errors.New("Getting Project Plans returned: " + response.Status)
	}

	return statusResponse, nil
}
