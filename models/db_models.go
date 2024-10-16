package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Connector struct {
	ID                     uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // gen_random_uuid() works with pg14+ source: https://stackoverflow.com/a/68539039/10269515
	CreatedAt              time.Time      `json:"-" gorm:"column:created_at"`
	UpdatedAt              time.Time      `json:"-" gorm:"column:updated_at"`
	DeletedAt              gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"`
	APIUrl                 string         `json:"api_url" gorm:"column:api_url"`
	APIKey                 string         `json:"-" gorm:"column:api_key"` // Hide APIKey in responses
	UpdatedBySubID         string         `json:"-" gorm:"column:updated_by_sub_id"`
	OrgID                  string         `json:"-" gorm:"column:org_id;unique"`
	AvailableToAllOrgUsers bool           `json:"-" gorm:"default:true"` // potential future column. in case the record belongs to single user instead of org
}
