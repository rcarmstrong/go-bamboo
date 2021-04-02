package bamboo

import (
	"fmt"
	"net/http"
)

// DeployService handles communication with the deploy related methods
type DeployService service

type IDeployService interface {
	CreateDeployVersion(deploymentProjectID int, planResultKey, versionName, nextVersionName string) (*DeployVersionResult, error)
	CreateDeploymentProject(deploymentProjectRequest CreateDeploymentProjectRequest) (dp DeploymentProject, err error)
	DeleteDeploymentProject(id uint) (result bool, err error)
	UpdateDeploymentProject(projectID uint, deploymentProjectRequest UpdateDeploymentProjectRequest) (dp DeploymentProject, err error)
	ListDeploys() (DeploysResponse, error)
	ListDeploysForPlan(planKey string) (DeploysResponse, error)
	DeployEnvironments(id int) (*DeployEnvironment, error)
	DeployEnvironmentResults(id int) (*DeployEnvironmentResults, error)
	QueueDeploy(environmentID, versionID int) (*QueueDeployRequest, error)
	DeployStatus(id int) (*DeployStatus, error)
	AddRepository(projectID uint, params AddRepositoryRequest) (r Repository, err error)
	DeleteRepository(projectID, repoID uint) (r Repository, err error)
}

type AddRepositoryRequest struct {
	ID uint `json:"id"`
}

// DeployResponse is the REST response from the server
type DeployResponse struct {
	*ResourceMetadata
}

// DeploysResponse is a collection of Deploy elements
type DeploysResponse = []*DeploymentProject

// Deploy is a single Deploy definition
type Deploy struct {
	ID           int                  `json:"id"`
	PlanKey      *PlanKey             `json:"planKey,omitempty"`
	Name         string               `json:"name,omitempty"`
	Description  string               `json:"description,omitempty"`
	Environments []*DeployEnvironment `json:"environments,omitempty"`
}

type DeploymentProject struct {
	ID                     uint                `json:"id"`
	Oid                    string              `json:"oid"`
	Key                    PlanKey             `json:"key"`
	Name                   string              `json:"name"`
	PlanKey                PlanKey             `json:"planKey"`
	Description            string              `json:"description"`
	Environments           []DeployEnvironment `json:"environments"`
	RepositorySpecsManaged bool                `json:"repositorySpecsManaged"`
}

type DeployEnvironment struct {
	ID                  uint                        `json:"id"`
	Key                 PlanKey                     `json:"key"`
	Name                string                      `json:"name"`
	Description         string                      `json:"description"`
	DeploymentProjectId uint                        `json:"deploymentProjectId"`
	Operations          DeployEnvironmentOperations `json:"operations"`
	Position            uint                        `json:"position"`
	ConfigurationState  string                      `json:"configurationState"`
}

type DeployEnvironmentOperations struct {
	CanView                   bool   `json:"canView"`
	CanEdit                   bool   `json:"canEdit "`
	CanDelete                 bool   `json:"canDelete"`
	AllowedToExecute          bool   `json:"allowedToExecute"`
	CanExecute                bool   `json:"canExecute"`
	CantExecuteReason         string `json:"cantExecuteReason"`
	AllowedToCreateVersion    bool   `json:"allowedToCreateVersion"`
	AllowedToSetVersionStatus bool   `json:"allowedToSetVersionStatus"`
}

// DeployEnvironmentResults is the information for a single Deploy
type DeployEnvironmentResults struct {
	Name    string          `json:"name"`
	ID      int             `json:"id"`
	Results []*DeployStatus `json:"results"`
}

// DeploymentVersion contains version information for a deployment
type DeploymentVersion struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// QueueDeployRequest contains information from a queue deploy request
type QueueDeployRequest struct {
	DeploymentResultID int   `json:"deploymentResultId"`
	Link               *Link `json:"link"`
}

// DeployStatus contains deploy status information
type DeployStatus struct {
	DeploymentVersion     *DeploymentVersion `json:"deploymentVersion"`
	DeploymentVersionName string             `json:"deploymentVersionName"`
	DeploymentState       string             `json:"deploymentState"`
	LifeCycleState        string             `json:"lifeCycleState"`
	StartedDate           int                `json:"startedDate"`
}

type createDeploymentVersion struct {
	PlanResultKey   string `json:"planResultKey"`
	Name            string `json:"name"`
	NextVersionName string `json:"nextVersionName"`
}

// DeployVersionResult will have the information for creating a
// new release/version for bamboo
type DeployVersionResult struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DeployVersionListResult stores a list of deployment versions
type DeployVersionListResult struct {
	Versions []*DeployVersionResult `json:"versions"`
}

type CreateDeploymentProjectRequest struct {
	Name         string  `json:"name"`
	PlanKey      PlanKey `json:"planKey"`
	Description  string  `json:"description"`
	PublicAccess bool    `json:"publicAccess"`
}

type UpdateDeploymentProjectRequest struct {
	Name        string  `json:"name"`
	PlanKey     PlanKey `json:"planKey"`
	Description string  `json:"description"`
}

// CreateDeployVersion will take a deploy project id, plan result, version name and the next version name and create a release.
func (d *DeployService) CreateDeployVersion(deploymentProjectID int, planResultKey, versionName, nextVersionName string) (*DeployVersionResult, error) {
	createDeployment := &createDeploymentVersion{
		PlanResultKey:   planResultKey,
		Name:            versionName,
		NextVersionName: nextVersionName,
	}

	request, err := d.client.NewRequest(http.MethodPost, fmt.Sprintf("deploy/project/%d/version", deploymentProjectID), createDeployment)
	if err != nil {
		return nil, err
	}

	result := &DeployVersionResult{}

	response, err := d.client.Do(request, result)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, newRespErr(response, "Error creating deploy version")
	}

	return result, nil
}

func (d *DeployService) DeleteDeploymentProject(id uint) (result bool, err error) {
	request, err := d.client.NewRequest(http.MethodDelete, fmt.Sprintf("deploy/project/%d", id), nil)
	if err != nil {
		return
	}

	response, err := d.client.Do(request, nil)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusNoContent {
		return result, newRespErr(response, "Error delete deployment project")
	}

	return true, err
}

func (d *DeployService) CreateDeploymentProject(deploymentProjectRequest CreateDeploymentProjectRequest) (dp DeploymentProject, err error) {
	request, err := d.client.NewRequest(http.MethodPut, "deploy/project", deploymentProjectRequest)
	if err != nil {
		return
	}

	response, err := d.client.Do(request, &dp)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		return dp, newRespErr(response, "Error create deployment project")
	}

	return
}

func (d *DeployService) UpdateDeploymentProject(projectID uint, deploymentProjectRequest UpdateDeploymentProjectRequest) (dp DeploymentProject, err error) {
	request, err := d.client.NewRequest(http.MethodPost, fmt.Sprintf("deploy/project/%d", projectID), deploymentProjectRequest)
	if err != nil {
		return
	}

	response, err := d.client.Do(request, &dp)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		return dp, newRespErr(response, "Error update deployment project")
	}

	return
}

func (d *DeployService) AddRepository(projectID uint, params AddRepositoryRequest) (r Repository, err error) {
	request, err := d.client.NewRequest(http.MethodPost, fmt.Sprintf("deploy/project/%d/repository", projectID), params)
	if err != nil {
		return
	}

	response, err := d.client.Do(request, &r)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusCreated {
		return r, newRespErr(response, "Error add repository to project")
	}

	return
}

func (d *DeployService) DeleteRepository(projectID, repoID uint) (r Repository, err error) {
	request, err := d.client.NewRequest(http.MethodDelete, fmt.Sprintf("deploy/project/%d/repository/%d", projectID, repoID), nil)
	if err != nil {
		return
	}

	response, err := d.client.Do(request, &r)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		return r, newRespErr(response, "Error delete repository from project")
	}

	return
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

	if response.StatusCode != http.StatusOK {
		return nil, newRespErr(response, "Error listing deploys")
	}

	return deployResp, nil
}

func (d *DeployService) ListDeploysForPlan(planKey string) (DeploysResponse, error) {
	request, err := d.client.NewRequest(http.MethodGet, fmt.Sprintf("deploy/project/forPlan?planKey=%s", planKey), nil)
	if err != nil {
		return nil, err
	}

	deployResp := DeploysResponse{}
	response, err := d.client.Do(request, &deployResp)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, newRespErr(response, "Error listing deploys")
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

	if response.StatusCode != http.StatusOK {
		return nil, newRespErr(response, "Error getting environments")
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

	if response.StatusCode != http.StatusOK {
		return nil, newRespErr(response, "Error getting deploy results")
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

	if response.StatusCode != http.StatusOK {
		return nil, newRespErr(response, "Error queueing deploy")
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

	if response.StatusCode != http.StatusOK {
		return nil, newRespErr(response, "Error getting deploy status")
	}

	return deployStatus, nil
}
