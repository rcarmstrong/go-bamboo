package bamboo

import (
	"fmt"
	"log"
	"net/http"
)

// Group contains information about a group of Bamboo users
type Group struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions,omitempty"`
}

type groupProjectPlanResponse struct {
	Results []Group `json:"results"`
}

// GroupPermissionsList returns a list of group permissions for the given resource. Leave Key blank when setting permissions globally.
func (p *Permissions) GroupPermissionsList(opts PermissionsOpts) ([]Group, *http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodGet, groupPermissionsListURL(opts.Resource, opts.Key), nil)
	if err != nil {
		return nil, nil, err
	}

	data := groupProjectPlanResponse{}
	response, err := p.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving group information for project %s returned %s", opts.Key, response.Status)}
	}

	return data.Results, response, nil
}

// GroupPermissions returns the group's permissions for the given resource. Leave Key blank when setting permissions globally.
func (p *Permissions) GroupPermissions(group string, opts PermissionsOpts) ([]string, *http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodGet, groupPermissionsURL(opts.Resource, opts.Key, group), nil)
	if err != nil {
		return nil, nil, err
	}

	data := groupProjectPlanResponse{}
	response, err := p.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving group information for project %s returned %s", opts.Key, response.Status)}
	}

	if len(data.Results) == 0 {
		response.StatusCode = http.StatusNoContent
		return nil, response, nil
	}

	return data.Results[0].Permissions, response, nil
}

// SetGroupPermissions sets the group's permissions for the given resource. Leave Key blank when setting permissions globally.
func (p *Permissions) SetGroupPermissions(group string, permissions []string, opts PermissionsOpts) (*http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodPut, editGroupPermissionsURL(opts.Resource, opts.Key, group), permissions)
	if err != nil {
		return nil, err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return response, err
	}

	switch response.StatusCode {
	case 400:
		return response, &simpleError{"Group doesn't exist or one of the requested permission isn't supported for the given endpoint."}
	case 401:
		return response, &simpleError{"You must be an admin to preform this action"}
	case 304:
		log.Println("Group already had requested permissions and permission state hasn't been changed.")
	case 204:
		log.Println("Group's permissions were granted.")
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return response, nil
}

// RemoveGroupPermissions removes the given permissions from the group's permissions for the given project's plans. Leave Key blank when setting permissions globally.
func (p *Permissions) RemoveGroupPermissions(group string, permissions []string, opts PermissionsOpts) (*http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodDelete, editGroupPermissionsURL(opts.Resource, opts.Key, group), permissions)
	if err != nil {
		return nil, err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return response, err
	}

	switch response.StatusCode {
	case 400:
		return response, &simpleError{"Group doesn't exist or one of the requested permission isn't supported for the given endpoint."}
	case 401:
		return response, &simpleError{"You must be an admin to preform this action"}
	case 304:
		log.Println("Group already lacked requested permissions and permission state hasn't been changed")
	case 204:
		log.Println("Group's permissions were revoked.")
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return response, nil
}

// AvailableGroupsPermissionsList returns a list of groups which weren't explicitly granted any permissions to the resource. Leave Key blank when setting permissions globally.
func (p *Permissions) AvailableGroupsPermissionsList(opts PermissionsOpts) ([]Group, *http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodGet, availableGroupsURL(opts.Resource, opts.Key), nil)
	if err != nil {
		return nil, nil, err
	}

	data := groupProjectPlanResponse{}
	response, err := p.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving group permission information for project %s returned %s", opts.Key, response.Status)}
	}

	return data.Results, response, nil
}
