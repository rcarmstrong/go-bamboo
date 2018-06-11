package bamboo

// WritePermission the sting the API expects for write permissions.
// Allows a user to view and edit the configuration of the plan and its jobs, not including permissions or stages.
const WritePermission string = "WRITE"

// BuildPermission the sting the API expects for build permissions.
// Allows a user to trigger a manual build, or suspend and resume the plan.
const BuildPermission string = "BUILD"

// ReadPermission the sting the API expects for read permissions.
// Allows a user to view the plan and its builds.
const ReadPermission string = "READ"

// ClonePermission the sting the API expects for clone permissions.
// Allows a user to clone the plan.
const ClonePermission string = "CLONE"

// AdminPermission is the sting the API expects for admin permissions.
// Allows a user to edit all aspects of the plan including permissions and stages.
const AdminPermission string = "ADMINISTRATION"

// Permissions is the container for all permissions related endpoints
type Permissions service
