package bamboo

var (
	permissionsTestCases = []PermissionsOpts{
		PermissionsOpts{
			Resource: GlobalResource,
		},
		PermissionsOpts{
			Resource: PlanResource,
			Key:      "TEST",
		},
		PermissionsOpts{
			Resource: RepositoryResource,
			Key:      "TEST",
		},
		PermissionsOpts{
			Resource: ProjectResource,
			Key:      "TEST",
		},
		PermissionsOpts{
			Resource: EnvironmentResource,
			Key:      "TEST",
		},
		PermissionsOpts{
			Resource: ProjectPlanResource,
			Key:      "TEST",
		},
		PermissionsOpts{
			Resource: DeploymentResource,
			Key:      "TEST",
		},
	}

	allowedPlanAccessLevels = [][]string{
		// Allowed to view a plan
		[]string{ReadPermission},
		// Allowed to view and edit a plan. Cannot manually start builds.
		[]string{ReadPermission, WritePermission},
		// Allowed to view and manually build a plan. Cannot edit the plan.
		[]string{ReadPermission, BuildPermission},
		// Allowed to view the plan and clone the plan for a new plan. Cannot edit or manually build the plan.
		[]string{ReadPermission, ClonePermission},
		// Allowed to view, edit, and build the plan. Cannot clone the plan.
		[]string{ReadPermission, WritePermission, BuildPermission},
		// Allowed to view, edit, and clone the plan for a new plan. Cannot manually start builds.
		[]string{ReadPermission, WritePermission, ClonePermission},
		// Allowed to view, build, and clone the plan. Cannot edit the plan.
		[]string{ReadPermission, BuildPermission, ClonePermission},
		// Allowed to view, edit, build, and clone the plan.
		[]string{ReadPermission, WritePermission, BuildPermission, ClonePermission},
		// Admin access to plan. Once admin access is granted, all other permissions are allowed and cannot be resticted.
		[]string{ReadPermission, WritePermission, BuildPermission, ClonePermission, AdminPermission},
	}

	allowedProjectAccessLevels = [][]string{
		// Allowed to create a plan in a given project
		[]string{CreatePermission},
		// Allowed to create and administer all plans in a project
		[]string{CreatePermission, AdminPermission},
	}
)
