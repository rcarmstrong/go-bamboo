package bamboo

import (
	"fmt"
)

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
		return fmt.Sprintf(permissionBase+"/groups?name=%s", GlobalResource, groupname)
	}
	return fmt.Sprintf(permissionBase+"/%s/groups?name=%s", resource, key, groupname)
}

func editGroupPermissionsURL(resource, key, groupname string) string {
	if resource == GlobalResource {
		return fmt.Sprintf(permissionBase+"/groups/%s", GlobalResource, groupname)
	}
	return fmt.Sprintf(permissionBase+"/%s/groups/%s", resource, key, groupname)
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
