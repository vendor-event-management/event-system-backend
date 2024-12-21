package domain

type RoleType string
type EventStatus string

const (
	// RoleType Enum of Users
	HR     RoleType = "HR"
	Vendor RoleType = "Vendor"

	// EventStatus Enum of Events
	Pending  EventStatus = "Pending"
	Approved EventStatus = "Approved"
	Rejected EventStatus = "Rejected"
)
