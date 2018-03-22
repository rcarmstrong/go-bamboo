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
	results []Group
}

// GroupPermissionsList returns a list of groups which have plan permissions for the given project with page limits set
// by Pagination.Start and Pagination.Limit. If Pagination is nil, then start is 0 and limit is 25.
func (pr *ProjectPlanService) GroupPermissionsList(projectKey string, pagination *Pagination) ([]Group, *http.Response, error) {
	if pagination == nil {
		pagination = &Pagination{
			Start: 0,
			Limit: 25,
		}
	}

	u := fmt.Sprintf("projectplan/%s/groups?start=%d&limit=%d", projectKey, pagination.Start, pagination.Limit)
	request, err := pr.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := groupProjectPlanResponse{}
	response, err := pr.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving group information for project %s returned %s", projectKey, response.Status)}
	}

	return data.results, nil, nil
}

// GroupPermissions returns the group's permissions for the given project.
func (pr *ProjectPlanService) GroupPermissions(projectKey, group string) ([]string, *http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/groups?name=%s", projectKey, group)
	request, err := pr.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := groupProjectPlanResponse{}
	response, err := pr.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving group information for project %s returned %s", projectKey, response.Status)}
	}

	if len(data.results) == 0 {
		return nil, nil, &simpleError{fmt.Sprintf("Group %s not found in project plan permissions for %s", group, projectKey)}
	}

	return data.results[0].Permissions, nil, nil
}

// SetGroupPermissions sets the group's permissions for the given project's plans to the passed in permissions array
func (pr *ProjectPlanService) SetGroupPermissions(projectKey, group string, permissions []string) (*http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/groups/%s", projectKey, group)
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
		return response, &simpleError{"Group doesn't exist or one of the requested permission isn't supported for the given endpoint."}
	case 401:
		return response, &simpleError{"You must be an admin to preform this action"}
	case 304:
		log.Println("Group already had requested permissions and permission state hasn't been changed.")
	case 204:
		log.Println("Group's permissions were granted.")
		return nil, nil
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// RemoveGroupPermissions removes the given permissions from the group's permissions for the given project's plans
func (pr *ProjectPlanService) RemoveGroupPermissions(projectKey, group string, permissions []string) (*http.Response, error) {
	u := fmt.Sprintf("projectplan/%s/group/%s", projectKey, group)
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
		return response, &simpleError{"Group doesn't exist or one of the requested permission isn't supported for the given endpoint."}
	case 401:
		return response, &simpleError{"You must be an admin to preform this action"}
	case 304:
		log.Println("Group already lacked requested permissions and permission state hasn't been changed")
	case 204:
		log.Println("Group's permissions were revoked.")
		return nil, nil
	default:
		return response, &simpleError{fmt.Sprintf("Server responded with unexpected return code %d", response.StatusCode)}
	}
	return nil, nil
}

// AvailableGroupPermissionsList return a list of groups which weren't explicitly granted any project plan permissions for the
// given project. Page limits are set by Pagination.Start and Pagination.Limit. If Pagination is nil, then start is 0 and limit is 25.
func (pr *ProjectPlanService) AvailableGroupPermissionsList(projectKey string, pagination *Pagination) ([]Group, *http.Response, error) {
	if pagination == nil {
		pagination = &Pagination{
			Start: 0,
			Limit: 25,
		}
	}

	u := fmt.Sprintf("projectplan/%s/available-groups?start=%d&limit=%d", projectKey, pagination.Start, pagination.Limit)
	request, err := pr.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := groupProjectPlanResponse{}
	response, err := pr.client.Do(request, &data)
	if err != nil {
		return nil, response, err
	}

	if response.StatusCode == 401 {
		return nil, response, &simpleError{"You must be an admin to access this information"}
	} else if response.StatusCode != 200 {
		return nil, response, &simpleError{fmt.Sprintf("Retrieving group permission information for project %s returned %s", projectKey, response.Status)}
	}

	return data.results, nil, nil
}
