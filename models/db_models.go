package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CredentialsType string

const (
	CredentialsTypeAPIKey            CredentialsType = "api_key"
	CredentialsTypeClientCredentials CredentialsType = "client_credentials"
	CredentialsTypePAT               CredentialsType = "pat" // personal access token
)

type Connector struct {
	ID                     uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // gen_random_uuid() works with pg14+ source: https://stackoverflow.com/a/68539039/10269515
	CreatedAt              time.Time       `json:"-" gorm:"column:created_at"`
	UpdatedAt              time.Time       `json:"-" gorm:"column:updated_at"`
	DeletedAt              gorm.DeletedAt  `json:"-" gorm:"column:deleted_at;index"`
	APIUrl                 string          `json:"api_url" gorm:"column:api_url"`
	Credentials            string          `json:"-" gorm:"column:credentials"` // Hide credentials in responses
	AuthTokenUrl           string          `json:"-" gorm:"column:auth_token_url"`
	CredentialsType        CredentialsType `json:"credentials_type" gorm:"column:credentials_type"`
	UpdatedBySubID         string          `json:"-" gorm:"column:updated_by_sub_id"`
	OrgID                  string          `json:"-" gorm:"column:org_id;unique"`
	AvailableToAllOrgUsers bool            `json:"-" gorm:"default:true"` // potential future column. in case the record belongs to single user instead of org
}

func (c *Connector) BeforeSave(tx *gorm.DB) (err error) {
	switch c.CredentialsType {
	case CredentialsTypeAPIKey, CredentialsTypeClientCredentials, CredentialsTypePAT:
		// Valid CredentialsType, do nothing
	default:
		// Invalid CredentialsType
		return fmt.Errorf("invalid credentials type: %s", c.CredentialsType)
	}
	return nil
}
