package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID            uuid.UUID      `gorm:"type:varchar(36);primary_key" json:"id"`
	Name          string         `gorm:"type:varchar(255)" json:"name"`
	PostalCode    string         `gorm:"type:varchar(10)" json:"postal_code"`
	Location      sql.NullString `gorm:"type:text" json:"location"`
	ProposedDates string         `gorm:"type:varchar(255)" json:"proposed_dates"`
	CreatedBy     uuid.UUID      `gorm:"type:uuid" json:"created_by"`
	CreatedAt     sql.NullTime   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     sql.NullTime   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     *time.Time     `gorm:"index" json:"deleted_at,omitempty"`
	CreatedByUser User           `gorm:"foreignkey:CreatedBy;association_foreignkey:ID" json:"created_by_user"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}

func (e *Event) TableName() string {
	return "events"
}
