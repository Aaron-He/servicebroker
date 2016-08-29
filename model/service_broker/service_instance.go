package model

//
type ServiceInstance struct {
	ID               string `json:"id"`
	DashboardURL     string `json:"dashboard_url"`
	InternalID       string `json:"internal_id, omitempty"`
	ServiceID        string `json:"service_id"`
	PlanID           string `json:"plan_id"`
	OrganizationGuid string `json:"organization_guid"`
	SpaceGuid        string `json:"space_guid"`

	LastOperation *LastOperation `json:"last_operation, omitempty"`

	Parameters interface{} `json:"parameters, omitempty"`
}

type LastOperation struct {
	State                    string `json:"state"`
	Description              string `json:"description"`
	AsyncPollIntervalSeconds int    `json:"async_poll_interval_seconds, omitempty"`
}

type ServiceInstanceRequest struct {
	OrgID             string                 `json:"organization_guid,omitempty"`
	PlanID            string                 `json:"plan_id,omitempty"`
	ServiceID         string                 `json:"service_id,omitempty"`
	SpaceID           string                 `json:"space_id,omitempty"`
	Parameters        map[string]interface{} `json:"parameters,omitempty"`
	AcceptsIncomplete bool                   `json:"accepts_incomplete,omitempty"`
}

type CreateServiceInstanceResponse struct {
	DashboardURL  string         `json:"dashboard_url, omitempty"`
	LastOperation *LastOperation `json:"last_operation, omitempty"`
}
