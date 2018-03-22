package bamboo

import (
	"fmt"
	"log"
	"net/http"
)

// AnonymousRole is the string the API expects for anonymous users
const AnonymousRole = "ANONYMOUS"

// LoggedInRole is the string the API expects for logged in users
const LoggedInRole = "LOGGED_IN"

type roleProjectPlanResponce struct {
	results []Role
}

// Role contains information about a role
type Role struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions,omitempty"`
}

// RolePermissionsList return a list of roles which have plan permissions for the given
// project. Currently, only Logged In Users and Anonymous Users roles are supported.
func (pr *ProjectPlanService) RolePermissionsList(projectKey string) ([]Role, *http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/roles", projectKey)
	request, err := pr.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := roleProjectPlanResponce{}
	response, err := pr.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving role information for project %s returned %s", projectKey, response.Status)}
	}

	return data.results, nil, nil
}

// SetLoggedInUserPermissions sets the logged in users role's permissions for the given project's plans to the passed in permissions
func (pr *ProjectPlanService) SetLoggedInUserPermissions(projectKey string, permissions []string) (*http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/roles/%s", projectKey, LoggedInRole)
	request, err := pr.client.NewRequest("PUT", u, permissions)
	if err != nil {
		return nil, err
	}

	response, err := pr.client.Do(request, nil)
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
		return nil, nil
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// RemoveLoggedInUsersPermissions removes the given permissions from the logged in users role's permissions for the given project's plans
func (pr *ProjectPlanService) RemoveLoggedInUsersPermissions(projectKey string, permissions []string) (*http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/roles/%s", projectKey, LoggedInRole)
	request, err := pr.client.NewRequest("DELETE", u, permissions)
	if err != nil {
		return nil, err
	}

	response, err := pr.client.Do(request, nil)
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
		return nil, nil
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// SetAnonymousReadPermission allows anonymous users to view plans
func (pr *ProjectPlanService) SetAnonymousReadPermission(projectKey string) (*http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/roles/%s", projectKey, AnonymousRole)
	request, err := pr.client.NewRequest("PUT", u, []string{ReadPermission})
	if err != nil {
		return nil, err
	}

	response, err := pr.client.Do(request, nil)
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
		return nil, nil
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// RemoveAnonymousReadPermission removes the ability for anonymous users to view plans
func (pr *ProjectPlanService) RemoveAnonymousReadPermission(projectKey string) (*http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/roles/%s", projectKey, AnonymousRole)
	request, err := pr.client.NewRequest("DELETE", u, []string{ReadPermission})
	if err != nil {
		return nil, err
	}

	response, err := pr.client.Do(request, nil)
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
		return nil, nil
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}
