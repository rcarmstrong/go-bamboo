package bamboo

import (
	"fmt"
	"log"
	"net/http"
)

// Role contains information about a role
type Role struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions,omitempty"`
}

type roleProjectPlanResponce struct {
	Results []Role `json:"results"`
}

// RolePermissionsList returns the list of permissions for the roles on the given entity in the given resource
func (p *Permissions) RolePermissionsList(opts PermissionsOpts) ([]Role, *http.Response, error) {
	request, err := p.client.NewRequest(http.MethodGet, rolePermissionsListURL(opts.Resource, opts.Key), nil)
	if err != nil {
		return nil, nil, err
	}

	data := roleProjectPlanResponce{}
	response, err := p.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving role information for project %s returned %s", opts.Key, response.Status)}
	}

	return data.Results, response, nil
}

// SetLoggedInUsersPermissions sets the logged in users role's permissions for the given project's plans to the passed in permissions
func (p *Permissions) SetLoggedInUsersPermissions(permissions []string, opts PermissionsOpts) (*http.Response, error) {
	request, err := p.client.NewRequest(http.MethodPut, loggedInRolePermissionsURL(opts.Resource, opts.Key), permissions)
	if err != nil {
		return nil, err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return response, err
	}

	switch response.StatusCode {
	case 401:
		return response, &simpleError{"You must be an admin to preform this action"}
	case 304:
		log.Println("Logged In Users Role already had requested permissions and permission state hasn't been changed.")
	case 204:
		log.Println("Logged In Users Role's permissions were granted.")
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// RemoveLoggedInUsersPermissions removes the given permissions from the logged in users role's permissions for the given project's plans
func (p *Permissions) RemoveLoggedInUsersPermissions(permissions []string, opts PermissionsOpts) (*http.Response, error) {
	request, err := p.client.NewRequest(http.MethodDelete, loggedInRolePermissionsURL(opts.Resource, opts.Key), permissions)
	if err != nil {
		return nil, err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return response, err
	}

	switch response.StatusCode {
	case 401:
		return response, &simpleError{"You must be an admin to preform this action"}
	case 304:
		log.Println("Logged In Users Role already lacked requested permissions and permission state hasn't been changed")
	case 204:
		log.Println("Logged In Users Role's permissions were revoked.")
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// SetAnonymousReadPermission allows anonymous users to view plans
func (p *Permissions) SetAnonymousReadPermission(opts PermissionsOpts) (*http.Response, error) {
	request, err := p.client.NewRequest(http.MethodPut, anonymousRolePermissionsURL(opts.Resource, opts.Key), []string{ReadPermission})
	if err != nil {
		return nil, err
	}

	response, err := p.client.Do(request, nil)
	if err != nil {
		return response, err
	}

	switch response.StatusCode {
	case 401:
		return response, &simpleError{"You must be an admin to preform this action"}
	case 304:
		log.Println("Anonymous Role already had requested permissions and permission state hasn't been changed.")
	case 204:
		log.Println("Anonymous Role's permissions were granted.")
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// RemoveAnonymousReadPermission removes the ability for anonymous users to view plans
func (p *Permissions) RemoveAnonymousReadPermission(opts PermissionsOpts) (*http.Response, error) {
	request, err := p.client.NewRequest(http.MethodDelete, anonymousRolePermissionsURL(opts.Resource, opts.Key), []string{ReadPermission})
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
		log.Println("Anonymous Role already lacked requested permissions and permission state hasn't been changed")
	case 204:
		log.Println("Anonymous Role's permissions were revoked.")
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}
