package urls

import (
	"fmt"
)

const permissionBase = "permissions/%s"

// PlanResource is the URL piece when getting plan permissions
const PlanResource string = "plan"

// GlobalResource is the URL piece when getting global permissions
const GlobalResource string = "global"

// GroupResource is the URL piece when getting group permissions
const GroupResource string = "group"

// RepositoryResource is the URL piece when getting repository permissions
const RepositoryResource string = "repository"

// ProjectResource is the URL piece when getting project permissions
const ProjectResource string = "project"

// EnvironmentResource is the URL piece when getting environment permissions
const EnvironmentResource string = "environment"

// ProjectPlanResource is the URL piece when getting projectplan permissions
const ProjectPlanResource string = "projectplan"

// DeploymentResource is the URL piece when getting deployment permissions
const DeploymentResource string = "deployment"

// -- Permissions --
// Users
func userPermissionsListURL(resource, key string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/users", GlobalResource)
	}
	return fmt.Sprintf(permissionBase+"/%s/users", resource, key)
}

func userPermissionsURL(resource, key, username string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/users?name=%s", GlobalResource, username)
	}
	return fmt.Sprintf(permissionBase+"/%s/users?name=%s", resource, key, username)
}

func editUserPermissionsURL(resource, key, username string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/users/%s", GlobalResource, username)
	}
	return fmt.Sprintf(permissionBase+"/%s/users/%s", resource, key, username)
}

func availableUsersURL(resource, key string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/available-users", GlobalResource)
	}
	return fmt.Sprintf(permissionBase+"/%s/available-users", resource, key)
}

// Groups
func groupPermissionsListURL(resource, key string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/groups", GlobalResource)
	}
	return fmt.Sprintf(permissionBase+"/%s/groups", resource, key)
}

func groupPermissionsURL(resource, key, groupname string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/group?name=%s", GlobalResource, groupname)
	}
	return fmt.Sprintf(permissionBase+"/%s/group?name=%s", resource, key, groupname)
}

func editGroupPermissionsURL(resource, key, groupname string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/group/%s", GlobalResource, groupname)
	}
	return fmt.Sprintf(permissionBase+"/%s/group/%s", resource, key, groupname)
}

func availableGroupsURL(resource, key string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/available-groups", GlobalResource)
	}
	return fmt.Sprintf(permissionBase+"/%s/available-groups", resource, key)
}

// Roles
func rolePermissionsListURL(resource, key string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/roles", GlobalResource)
	}
	return fmt.Sprintf(permissionBase+"/%s/roles", resource, key)
}

func loggedInRolePermissionsURL(resource, key string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/roles/LOGGED_IN", GlobalResource)
	}
	return fmt.Sprintf(permissionBase+"/%s/roles/LOGGED_IN", resource, key)
}

func anonymousRolePermissionsURL(resource, key string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/roles/ANONYMOUS", GlobalResource)
	}
	return fmt.Sprintf(permissionBase+"/%s/roles/ANONYMOUS", resource, key)
}
