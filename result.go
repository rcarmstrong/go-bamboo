package bamboo

import (
	"fmt"
	"net/http"
)

// ResultService handles communication with build results
type ResultService service

// Result represents all the information associated with a build result
type Result struct {
	ChangeSet              `json:"changes"`
	ID                     int    `json:"id"`
	PlanName               string `json:"planName"`
	ProjectName            string `json:"projectName"`
	BuildResultKey         string `json:"buildResultKey"`
	LifeCycleState         string `json:"lifeCycleState"`
	BuildStartedTime       string `json:"buildStartedTime"`
	BuildCompletedTime     string `json:"buildCompletedTime"`
	BuildDurationInSeconds int    `json:"buildDurationInSeconds"`
	VcsRevisionKey         string `json:"vcsRevisionKey"`
	BuildTestSummary       string `json:"buildTestSummary"`
	SuccessfulTestCount    int    `json:"successfulTestCount"`
	FailedTestCount        int    `json:"failedTestCount"`
	QuarantinedTestCount   int    `json:"quarantinedTestCount"`
	SkippedTestCount       int    `json:"skippedTestCount"`
	Finished               bool   `json:"finished"`
	Successful             bool   `json:"successful"`
	BuildReason            string `json:"buildReason"`
	ReasonSummary          string `json:"reasonSummary"`
	Key                    string `json:"key"`
	State                  string `json:"state"`
	BuildState             string `json:"buildState"`
	Number                 int    `json:"number"`
	BuildNumber            int    `json:"buildNumber"`
}

// ChangeSet represents a collection of type Change
type ChangeSet struct {
	Set []Change `json:"change"`
}

// Change represents the author and commit hash of a source code change
type Change struct {
	Author      string `json:"author"`
	ChangeSetID string `json:"changesetId"`
}

// LatestResult returns the latest result information for the given plan key
func (r *ResultService) LatestResult(key string) (*Result, *http.Response, error) {
	request, err := r.client.NewRequest(http.MethodGet, latestResultURL(key), nil)
	if err != nil {
		return nil, nil, err
	}

	// query := request.URL.Query()
	// query.Add("expand", "changes,metadata,plan,vcsRevisions,artifacts,comments,labels,jiraIssues,stages")
	// request.URL.RawQuery = query.Encode()

	result := Result{}
	response, err := r.client.Do(request, &result)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("API returned unexpected status code %d", response.StatusCode)}
	}

	return &result, response, err
}
