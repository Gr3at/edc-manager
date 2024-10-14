package models

type Connector struct {
	APIUrl string `json:"api_url"`
	APIKey string `json:"-"` // Hide APIKey in responses
	// SubID                  string `json:"sub_id"`
	// OrgID                  string `json:"org_id"`
	AvailableToAllOrgUsers bool `json:"available_to_all_org_users"`
}
