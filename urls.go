package bamboo

import (
	"fmt"
)

// Permissions Endpoints
const permissionBase = "permissions/%s"

// Users
func userPermissionsListURL(resource, key string) string {
	return fmt.Sprintf(permissionBase+"/%s/users", resource, key)
}

func userPermissionsURL(resource, key, username string) string {
	return fmt.Sprintf(permissionBase+"/%s/users?name=%s", resource, key, username)
}

func editUserPermissionsURL(resource, key, username string) string {
	return fmt.Sprintf(permissionBase+"/%s/users/%s", resource, key, username)
}

func availableUsersURL(resource, key string) string {
	return fmt.Sprintf(permissionBase+"/%s/available-users", resource, key)
}

// Groups
func groupPermissionsListURL(resource, key string) string {
	return fmt.Sprintf(permissionBase+"/%s/groups", resource, key)
}

func groupPermissionsURL(resource, key, groupname string) string {
	return fmt.Sprintf(permissionBase+"/%s/group?name=%s", resource, key, groupname)
}

func editGroupPermissionsURL(resource, key, groupname string) string {
	return fmt.Sprintf(permissionBase+"/%s/group/%s", resource, key, groupname)
}

func availableGroupsURL(resource, key string) string {
	return fmt.Sprintf(permissionBase+"/%s/available-groups", resource, key)
}

// Roles
func rolePermissionsListURL(resource, key string) string {
	return fmt.Sprintf(permissionBase+"/%s/roles", resource, key)
}

func loggedInRolePermissionsURL(resource, key string) string {
	return fmt.Sprintf(permissionBase+"/%s/roles/LOGGED_IN", resource, key)
}

func anonymousRolePermissionsURL(resource, key string) string {
	return fmt.Sprintf(permissionBase+"/%s/roles/ANONYMOUS", resource, key)
}
