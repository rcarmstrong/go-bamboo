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

// UserProjectPlanResponse is the result of project plan user information request
type UserProjectPlanResponse struct {
	Results []User `json:"results"`
}

// UserPermissionsList returns a list of users which have plan permissions for the given project with page
// limits set by Pagination.Start and Pagination.Limit. If Pagination is nil, then start is 0 and limit is 25.
func (pr *ProjectPlanService) UserPermissionsList(projectKey string, pagination *Pagination) ([]User, *http.Response, error) {
	if pagination == nil {
		pagination = &Pagination{
			Start: 0,
			Limit: 25,
		}
	}

	u := fmt.Sprintf("projectplan/%s/users?start=%d&limit=%d", projectKey, pagination.Start, pagination.Limit)
	request, err := pr.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := UserProjectPlanResponse{}
	response, err := pr.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving user information for project %s returned %s", projectKey, response.Status)}
	}

	return data.Results, nil, nil
}

// UserPermissions returns the user permissions for the given user for the given project.
func (pr *ProjectPlanService) UserPermissions(projectKey, username string) ([]string, *http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/users?name=%s", projectKey, username)
	request, err := pr.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := UserProjectPlanResponse{}
	response, err := pr.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving user information for project %s returned %s", projectKey, response.Status)}
	}

	if len(data.Results) == 0 {
		return nil, response, &simpleError{fmt.Sprintf("User %s not found in project plan permissions for %s", username, projectKey)}
	}

	return data.Results[0].Permissions, nil, nil
}

// SetUserPermissions sets the users permissions for the given project's plans to the passed in permissions array
func (pr *ProjectPlanService) SetUserPermissions(projectKey, username string, permissions []string) (*http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/users/%s", projectKey, username)
	request, err := pr.client.NewRequest("PUT", u, permissions)
	if err != nil {
		return nil, err
	}

	response, err := pr.client.Do(request, nil)
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
		return nil, nil
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// RemoveUserPermissions removes the given permissions from the users permissions for the given project's plans
func (pr *ProjectPlanService) RemoveUserPermissions(projectKey, username string, permissions []string) (*http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/users/%s", projectKey, username)
	request, err := pr.client.NewRequest("DELETE", u, permissions)
	if err != nil {
		return nil, err
	}

	response, err := pr.client.Do(request, nil)
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
		return nil, nil
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// AvailableUserPermissionsList return a list of users which weren't explicitly granted any project plan permissions for the
// given project. Page limits are set by Pagination.Start and Pagination.Limit. If Pagination is nil, then start is 0 and limit is 25.
func (pr *ProjectPlanService) AvailableUserPermissionsList(projectKey string, pagination *Pagination) ([]User, *http.Response, error) {
	if pagination == nil {
		pagination = &Pagination{
			Start: 0,
			Limit: 25,
		}
	}

	u := fmt.Sprintf("projectplan/%s/available-users?start=%d&limit=%d", projectKey, pagination.Start, pagination.Limit)
	request, err := pr.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := UserProjectPlanResponse{}
	response, err := pr.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving user information for project %s returned %s", projectKey, response.Status)}
	}

	return data.Results, nil, nil
}
