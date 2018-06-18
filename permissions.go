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

// CreatePermission is the string the API expects when allowing a user/group to create a resource
const CreatePermission string = "CREATE"

// CreateRepositoryPermission is the string the API expects when allowing a user/group to create a repository
const CreateRepositoryPermission string = "CREATEREPOSITORY"

// PlanResource is the URL piece when getting plan permissions
const PlanResource string = "plan"

// GlobalResource is the URL piece when getting global permissions
const GlobalResource string = "global"

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

var knownResources = map[string]bool{
	PlanResource:        true,
	GlobalResource:      true,
	RepositoryResource:  true,
	ProjectResource:     true,
	EnvironmentResource: true,
	ProjectPlanResource: true,
	DeploymentResource:  true,
}

// Permissions is the container for all permissions related endpoints
type Permissions service

// PermissionsOpts holds the name of the resource that permissions are being retrieved
// from and the key for the specific object in that resource.
// -- LEAVE KEY BLANK FOR GLOBAL PERMISSIONS --
type PermissionsOpts struct {
	Resource string
	Key      string
}
