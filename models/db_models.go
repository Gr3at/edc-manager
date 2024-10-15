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
	APIUrl                 string         `json:"api_url" gorm:"unique"`
	APIKey                 string         `json:"-" gorm:"column:api_key"` // Hide APIKey in responses
	SubID                  string         `json:"-" gorm:"column:sub_id"`
	OrgID                  string         `json:"-" gorm:"column:org_id"`
	AvailableToAllOrgUsers bool           `json:"-"` // `json:"available_to_all_org_users"`
}

// BeforeCreate is a GORM hook that runs before a record is inserted in the database
// It ensures that a UUID is generated if not provided
// func (c *Connector) BeforeCreate(tx *gorm.DB) (err error) {
// 	if c.ID == uuid.Nil {
// 		c.ID = uuid.New() // Generate a new UUID
// 	}
// 	return
// }
