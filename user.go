package bamboo

import (
	"fmt"
	"log"
	"net/http"
)

// User contains information about a Bamboo user account
type User struct {
	Name        string   `json:"name"`
	FullName    string   `json:"fullName"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions,omitempty"`
}

type userPermissionsResponse struct {
	Results []User `json:"results"`
}

// UserPermissionsList returns a list of users and their permissions for the given resource key in the service
func (p *Permissions) UserPermissionsList(opts PermissionsOpts) ([]User, *http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodGet, userPermissionsListURL(opts.Resource, opts.Key), nil)
	if err != nil {
		return nil, nil, err
	}

	data := userPermissionsResponse{}
	response, err := p.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving user information for resource %s in service %s returned %s", opts.Key, opts.Resource, response.Status)}
	}

	return data.Results, response, nil
}

// UserPermissions returns the permissions for the specified user on the given resource in the given service
func (p *Permissions) UserPermissions(username string, opts PermissionsOpts) (*User, *http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodGet, userPermissionsURL(opts.Resource, opts.Key, username), nil)
	if err != nil {
		return nil, nil, err
	}

	data := userPermissionsResponse{}
	response, err := p.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving user information for resource %s in service %s returned %s", opts.Key, opts.Resource, response.Status)}
	}

	if len(data.Results) == 0 {
		response.StatusCode = http.StatusNoContent
		return nil, response, nil
	}

	return &data.Results[0], response, nil
}

// SetUserPermissions sets the users permissions for the given project's plans to the passed in permissions array
func (p *Permissions) SetUserPermissions(username string, permissions []string, opts PermissionsOpts) (*http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodPut, editUserPermissionsURL(opts.Resource, opts.Key, username), permissions)
	if err != nil {
		return nil, err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return response, err
	}

	switch response.StatusCode {
	case 400:
		return response, &simpleError{"User doesn't exist or one of the requested permission isn't supported for the given endpoint."}
	case 401:
		return response, &simpleError{"You must be an admin to access this information"}
	case 304:
		log.Println("User already had requested permissions and permission state hasn't been changed.")
	case 204:
		log.Println("User's permissions were granted.")
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return response, nil
}

// RemoveUserPermissions removes the given permissions from the users permissions for the given project's plans
func (p *Permissions) RemoveUserPermissions(username string, permissions []string, opts PermissionsOpts) (*http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodDelete, editUserPermissionsURL(opts.Resource, opts.Key, username), permissions)
	if err != nil {
		return nil, err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return response, err
	}

	switch response.StatusCode {
	case 400:
		return response, &simpleError{"User doesn't exist or one of the requested permission isn't supported for the given endpoint."}
	case 401:
		return response, &simpleError{"You must be an admin to access this information"}
	case 304:
		log.Println("User already lacked requested permissions and permission state hasn't been changed")
	case 204:
		log.Println("User's permissions were revoked.")
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return response, nil
}

// AvailableUsersPermissionsList return a list of users which weren't explicitly granted any project plan permissions for the given project.
func (p *Permissions) AvailableUsersPermissionsList(opts PermissionsOpts) ([]User, *http.Response, error) {
	if !knownResources[opts.Resource] {
		return nil, nil, &simpleError{fmt.Sprintf("Unknown resource %s", opts.Resource)}
	}

	request, err := p.client.NewRequest(http.MethodGet, availableUsersURL(opts.Resource, opts.Key), nil)
	if err != nil {
		return nil, nil, err
	}

	data := userPermissionsResponse{}
	response, err := p.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving user information for project %s returned %s", opts.Key, response.Status)}
	}

	return data.Results, response, nil
}
