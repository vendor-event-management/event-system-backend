package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventApproval struct {
	ID            uuid.UUID      `gorm:"type:varchar(36);primary_key" json:"id"`
	EventID       uuid.UUID      `gorm:"type:varchar(36);" json:"event_id"`
	Status        EventStatus    `gorm:"type:enum('Pending', 'Approved', 'Rejected')" json:"status"`
	VendorID      uuid.UUID      `gorm:"type:varchar(36);" json:"vendor_id"`
	ConfirmedDate sql.NullTime   `gorm:"type:date;" json:"confirmed_date"`
	Remarks       sql.NullString `gorm:"type:text;" json:"remarks"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     *time.Time     `gorm:"index" json:"deleted_at,omitempty"`
	Event         Event          `gorm:"foreignkey:EventID;association_foreignkey:ID" json:"event"`
	Vendor        User           `gorm:"foreignkey:VendorID;association_foreignkey:ID" json:"vendor"`
}

func (e *EventApproval) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}

func (ea *EventApproval) TableName() string {
	return "event_approvals"
}
