package test

import "github.com/lotos2512/bamboo"

var (
	permissionsTestCases = []bamboo.PermissionsOpts{
		bamboo.PermissionsOpts{
			Resource: bamboo.GlobalResource,
		},
		{
			Resource: bamboo.PlanResource,
			Key:      "TEST",
		},
		bamboo.PermissionsOpts{
			Resource: bamboo.RepositoryResource,
			Key:      "TEST",
		},
		bamboo.PermissionsOpts{
			Resource: bamboo.ProjectResource,
			Key:      "TEST",
		},
		bamboo.PermissionsOpts{
			Resource: bamboo.EnvironmentResource,
			Key:      "TEST",
		},
		bamboo.PermissionsOpts{
			Resource: bamboo.ProjectPlanResource,
			Key:      "TEST",
		},
		{
			Resource: bamboo.DeploymentResource,
			Key:      "TEST",
		},
	}

	allowedPlanAccessLevels = [][]string{
		// Allowed to view a plan
		{bamboo.ReadPermission},
		// Allowed to view and edit a plan. Cannot manually start builds.
		{bamboo.ReadPermission, bamboo.WritePermission},
		// Allowed to view and manually build a plan. Cannot edit the plan.
		[]string{bamboo.ReadPermission, bamboo.BuildPermission},
		// Allowed to view the plan and clone the plan for a new plan. Cannot edit or manually build the plan.
		[]string{bamboo.ReadPermission, bamboo.ClonePermission},
		// Allowed to view, edit, and build the plan. Cannot clone the plan.
		[]string{bamboo.ReadPermission, bamboo.WritePermission, bamboo.BuildPermission},
		// Allowed to view, edit, and clone the plan for a new plan. Cannot manually start builds.
		[]string{bamboo.ReadPermission, bamboo.WritePermission, bamboo.ClonePermission},
		// Allowed to view, build, and clone the plan. Cannot edit the plan.
		[]string{bamboo.ReadPermission, bamboo.BuildPermission, bamboo.ClonePermission},
		// Allowed to view, edit, build, and clone the plan.
		[]string{bamboo.ReadPermission, bamboo.WritePermission, bamboo.BuildPermission, bamboo.ClonePermission},
		// Admin access to plan. Once admin access is granted, all other permissions are allowed and cannot be resticted.
		[]string{bamboo.ReadPermission, bamboo.WritePermission, bamboo.BuildPermission, bamboo.ClonePermission, bamboo.AdminPermission},
	}

	allowedProjectAccessLevels = [][]string{
		// Allowed to create a plan in a given project
		[]string{bamboo.CreatePermission},
		// Allowed to create and administer all plans in a project
		[]string{bamboo.CreatePermission, bamboo.AdminPermission},
	}
)
