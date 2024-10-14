package models

type ConnectorInput struct {
	APIUrl string `json:"api_url" binding:"required"`
	APIKey string `json:"api_key" binding:"required"`
	// SubID                  string `json:"sub_id" binding:"required"`
	// OrgID                  string `json:"org_id" binding:"required"`
	AvailableToAllOrgUsers bool `json:"available_to_all_org_users"`
}
