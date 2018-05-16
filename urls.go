package bamboo

// Project Plan Permission Endpoints
const projectPlanPermissionBase = "permissions/projectplan"

// Users
const projectPlanUserPermissionList = projectPlanPermissionBase + "/%s/users"
const projectPlanSpecificUserPermissions = projectPlanUserPermissionList + "?name=%s"
const projectPlanEditUserPermissions = projectPlanUserPermissionList + "/%s"
const projectPlanAvailableUsers = projectPlanPermissionBase + "/%s/available-users"

// Groups

// Roles
