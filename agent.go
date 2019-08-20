package bamboo

import (
	"fmt"
	"net/http"
	"strconv"
)

// ExecutorTypeAgent is an accepted executor type string for Local/Remote Agents
const ExecutorTypeAgent string = "AGENT"

// ExecutorTypeImage is an accepted executor type string for Elastic Agents
const ExecutorTypeImage string = "IMAGE"

// AssignmentTypeProject is an accepted assignment type string for projects
const AssignmentTypeProject string = "PROJECT"

// AssignmentTypePlan is an accepted assignment type string for plans
const AssignmentTypePlan string = "PLAN"

// AssignmentTypeJob is an accepted assignment type string for jobs
const AssignmentTypeJob string = "JOB"

// AssignmentTypeEnvironment is an accepted assignment type string for deploy environments
const AssignmentTypeEnvironment string = "ENVIRONMENT"

var knownExecutorTypes = map[string]bool{
	ExecutorTypeAgent: true,
	ExecutorTypeImage: true,
}

var knownAssignmentTypes = map[string]bool{
	AssignmentTypeEnvironment: true,
	AssignmentTypeJob:         true,
	AssignmentTypePlan:        true,
	AssignmentTypeProject:     true,
}

// AgentService handles communication with the agent related endpoints
type AgentService service

// AgentAssignmentInformation encapsulates the information returned
// from the agent assignment endpoint
type AgentAssignmentInformation struct {
	NameElements        []string `json:"nameElements,omitempty"`
	Description         string   `json:"description,omitempty"`
	ExecutableType      string   `json:"executableType,omitempty"`
	ExecutableID        int64    `json:"executableId,omitempty"`
	ExecutableTypeLabel string   `json:"executableTypeLabel,omitempty"`
	ExecutorType        string   `json:"executorType,omitempty"`
	ExecutorID          int64    `json:"executorId,omitempty"`
	CapabilitiesMatch   bool     `json:"capabilitiesMatch,omitempty"`
}

// GetAgentAssignments gets the current assignment information for the given executor (agent)
func (a *AgentService) GetAgentAssignments(executorType string, executorID int64) ([]*AgentAssignmentInformation, *http.Response, error) {
	var u string
	if !emptyStrings(executorType) {
		if knownExecutorTypes[executorType] {
			u = "agent/assignment"
		} else {
			return nil, nil, &simpleError{fmt.Sprintf("Unknown ExecutorType: %s", executorType)}
		}
	} else {
		return nil, nil, &simpleError{fmt.Sprintf("ExecutorType cannot be an empty string")}
	}

	request, err := a.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	values := request.URL.Query()
	values.Set("executorType", executorType)
	values.Set("executorId", strconv.FormatInt(executorID, 10))
	request.URL.RawQuery = values.Encode()

	agentAssignmentInformation := []*AgentAssignmentInformation{}
	response, err := a.client.Do(request, &agentAssignmentInformation)
	if err != nil {
		return nil, nil, err
	}

	return agentAssignmentInformation, response, nil
}

// DedicateAgent assigns the given executor (agent) to the given entity
func (a *AgentService) DedicateAgent(executorType, assignmentType string, executorID, entityID int64) (*AgentAssignmentInformation, *http.Response, error) {
	var u string
	if !emptyStrings(executorType, assignmentType) {
		if knownExecutorTypes[executorType] {
			if knownAssignmentTypes[assignmentType] {
				u = "agent/assignment"
			} else {
				return nil, nil, &simpleError{fmt.Sprintf("Unknown AssignmentType: %s", assignmentType)}
			}
		} else {
			return nil, nil, &simpleError{fmt.Sprintf("Unknown ExecutorType: %s", executorType)}
		}
	} else {
		return nil, nil, &simpleError{fmt.Sprintf("ExecutorType or AssignmentType cannot be an empty string")}
	}

	request, err := a.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	values := request.URL.Query()
	values.Set("executorType", executorType)
	values.Set("executorId", strconv.FormatInt(executorID, 10))
	values.Set("assignmentType", assignmentType)
	values.Set("entityId", strconv.FormatInt(entityID, 10))
	request.URL.RawQuery = values.Encode()

	agentAssignmentInformation := AgentAssignmentInformation{}
	response, err := a.client.Do(request, &agentAssignmentInformation)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, response, &simpleError{fmt.Sprintf("DedicateAgent call returned: %d", response.StatusCode)}
	}

	return &agentAssignmentInformation, response, nil
}

// RemoveAgentAssignment removes the specified executor (agent) assignmnt from the given entity
func (a *AgentService) RemoveAgentAssignment(executorType, assignmentType string, executorID, entityID int64) (*AgentAssignmentInformation, *http.Response, error) {
	var u string
	if !emptyStrings(executorType, assignmentType) {
		if knownExecutorTypes[executorType] {
			if knownAssignmentTypes[assignmentType] {
				u = "agent/assignment"
			} else {
				return nil, nil, &simpleError{fmt.Sprintf("Unknown AssignmentType: %s", assignmentType)}
			}
		} else {
			return nil, nil, &simpleError{fmt.Sprintf("Unknown ExecutorType: %s", executorType)}
		}
	} else {
		return nil, nil, &simpleError{fmt.Sprintf("ExecutorType or AssignmentType cannot be an empty string")}
	}

	request, err := a.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	values := request.URL.Query()
	values.Set("executorType", executorType)
	values.Set("executorId", strconv.FormatInt(executorID, 10))
	values.Set("assignmentType", assignmentType)
	values.Set("entityId", strconv.FormatInt(entityID, 10))
	request.URL.RawQuery = values.Encode()

	agentAssignmentInformation := AgentAssignmentInformation{}
	response, err := a.client.Do(request, &agentAssignmentInformation)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode != http.StatusNoContent {
		return nil, response, &simpleError{fmt.Sprintf("DedicateAgent call returned: %d", response.StatusCode)}
	}

	return &agentAssignmentInformation, response, nil
}
