package bamboo

import (
	"fmt"
	"net/http"
)

// DeployService handles communication with the deploy related methods
type DeployService service

// DeployResponse is the REST response from the server
type DeployResponse struct {
	*ResourceMetadata
}

// DeploysResponse is a collection of Deploy elements
type DeploysResponse = []*Deploy

// Deploy is a single Deploy definition
type Deploy struct {
	ID           int                  `json:"id"`
	PlanKey      *PlanKey             `json:"planKey,omitempty"`
	Name         string               `json:"name,omitempty"`
	Description  string               `json:"description,omitempty"`
	Environments []*DeployEnvironment `json:"environments,omitempty"`
}

// DeployEnvironment is the information for an environment
type DeployEnvironment struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description,omitempty"`
	DeploymentProjectID int    `json:"deploymentProjectId,omitempty"`
}

// DeployEnvironmentResults is the information for a single Deploy
type DeployEnvironmentResults struct {
	Name    string          `json:"name"`
	ID      int             `json:"id"`
	Results []*DeployStatus `json:"results"`
}

type DeploymentVersion struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type QueueDeployRequest struct {
	DeploymentResultID int   `json:"deploymentResultId"`
	Link               *Link `json:"link"`
}

type DeployStatus struct {
	DeploymentVersion     *DeploymentVersion `json:"deploymentVersion"`
	DeploymentVersionName string             `json:"deploymentVersionName"`
	DeploymentState       string             `json:"deploymentState"`
	LifeCycleState        string             `json:"lifeCycleState"`
	StartedDate           int                `json:"startedDate"`
}

// ListDeploys lists all deployments
func (d *DeployService) ListDeploys() (DeploysResponse, error) {
	request, err := d.client.NewRequest(http.MethodGet, "deploy/project/all", nil)
	if err != nil {
		return nil, err
	}

	deployResp := DeploysResponse{}
	response, err := d.client.Do(request, &deployResp)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, err
	}

	return deployResp, nil
}

// DeployEnvironments returns information on the requested environment
func (d *DeployService) DeployEnvironments(id int) (*DeployEnvironment, error) {
	request, err := d.client.NewRequest(http.MethodGet, fmt.Sprintf("deploy/project/%d", id), nil)
	if err != nil {
		return nil, err
	}

	deployEnvironmentResp := &DeployEnvironment{}
	response, err := d.client.Do(request, &deployEnvironmentResp)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, err
	}

	return deployEnvironmentResp, nil
}

// DeployEnvironmentResults returns result information for the requested environment
func (d *DeployService) DeployEnvironmentResults(id int) (*DeployEnvironmentResults, error) {
	request, err := d.client.NewRequest(http.MethodGet, fmt.Sprintf("deploy/environment/%d/results", id), nil)
	if err != nil {
		return nil, err
	}

	deployEnvironmentResultsResp := &DeployEnvironmentResults{}
	response, err := d.client.Do(request, &deployEnvironmentResultsResp)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, err
	}

	return deployEnvironmentResultsResp, nil
}

// QueueDeploy adds a deploy of the specified version to the given environment.
func (d *DeployService) QueueDeploy(environmentID, versionID int) (*QueueDeployRequest, error) {
	request, err := d.client.NewRequest(http.MethodPost, fmt.Sprintf("queue/deployment/?environmentId=%d&versionId=%d", environmentID, versionID), nil)
	if err != nil {
		return nil, err
	}

	queueDeployRequest := &QueueDeployRequest{}
	response, err := d.client.Do(request, &queueDeployRequest)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, err
	}

	return queueDeployRequest, nil
}

// DeployStatus returns information on the requested deploy
func (d *DeployService) DeployStatus(id int) (*DeployStatus, error) {
	request, err := d.client.NewRequest(http.MethodGet, fmt.Sprintf("deploy/result/%d", id), nil)
	if err != nil {
		return nil, err
	}

	deployStatus := &DeployStatus{}
	response, err := d.client.Do(request, &deployStatus)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, err
	}

	return deployStatus, nil
}
