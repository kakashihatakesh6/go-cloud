package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	UserTypeIndividual = "individual"
	UserTypeEnterprise = "enterprise"
)

type User struct {
	ID           string        `gorm:"type:uuid;default:gen_random_uuid()"`
	Fullname     string        `gorm:"type:varchar(255);not null" json:"name"`
	Username     string        `gorm:"type:varchar(255);uniqueIndex:idx_username,where:username != ''" json:"username"`
	UserType     string        `gorm:"type:varchar(50);not null" json:"user_type"`
	Email        string        `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password     string        `gorm:"type:varchar(255);not null" json:"password"`
	Role         string        `gorm:"type:varchar(50);not null" json:"role"`
	Organization Organizations `gorm:"foreignKey:OrganizationID" json:"-"`
	Status       string        `gorm:"type:varchar(20);default:'active'" json:"status"`
	Designation  string        `gorm:"type:varchar(255)" json:"designation"`
	Nationality  string        `gorm:"type:varchar(255)" json:"nationality"`

	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
