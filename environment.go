package bamboo

type Environment struct {
	ID                  uint                  `json:"id"`
	Key                 PlanKey               `json:"key"`
	Name                string                `json:"name"`
	Description         string                `json:"description"`
	DeploymentProjectId uint                  `json:"deploymentProjectId"`
	Operations          EnvironmentOperations `json:"operations"`
	Position            uint                  `json:"position"`
	ConfigurationState  string                `json:"configurationState"`
}

type EnvironmentOperations struct {
	CanView                   bool   `json:"canView"`
	CanEdit                   bool   `json:"canEdit "`
	CanDelete                 bool   `json:"canDelete"`
	AllowedToExecute          bool   `json:"allowedToExecute"`
	CanExecute                bool   `json:"canExecute"`
	CantExecuteReason         string `json:"cantExecuteReason"`
	AllowedToCreateVersion    bool   `json:"allowedToCreateVersion"`
	AllowedToSetVersionStatus bool   `json:"allowedToSetVersionStatus"`
}
