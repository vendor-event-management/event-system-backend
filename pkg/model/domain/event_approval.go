package domain

import (
	"time"

	"github.com/google/uuid"
)

type EventApproval struct {
	ID        uuid.UUID   `gorm:"type:varchar(36);primary_key" json:"id"`
	EventID   uuid.UUID   `gorm:"type:varchar(36);" json:"event_id"`
	Status    EventStatus `gorm:"type:enum('Pending', 'Approved', 'Rejected')" json:"status"`
	VendorID  uuid.UUID   `gorm:"type:varchar(36);" json:"vendor_id"`
	CreatedAt time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt time.Time   `gorm:"index" json:"deleted_at,omitempty"`
	Event     Event       `gorm:"foreignkey:EventID;association_foreignkey:ID" json:"event"`
	Vendor    User        `gorm:"foreignkey:VendorID;association_foreignkey:ID" json:"vendor"`
}

func (ea *EventApproval) TableName() string {
	return "event_approvals"
}
