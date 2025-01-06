package models

import (
	"time"

	"gorm.io/gorm"
)

type Organizations struct {
	ID             string         `gorm:"type:uuid;default:gen_random_uuid()"`
	Name           string         `gorm:"type:varchar(255);not null" json:"organization_name"`
	BillingDetails string         `gorm:"type:text" json:"billing_details"`
	CreatedAt      time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
