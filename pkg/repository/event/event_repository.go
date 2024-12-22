package repository

import (
	"event-system-backend/pkg/model/domain"
	"event-system-backend/pkg/model/dto/response"
	"time"
)

type EventRepository interface {
	CreateEvent(event domain.Event, eventApproval domain.EventApproval) error
	FindEventByID(eventID string) (domain.Event, error)
	FindAllEventsByUserInvolved(userId string, role domain.RoleType, page, size int, nameEvent, status string) ([]response.EventScanResponse, int64, error)
	FindDetailEventByID(id string) (response.EventDetailScanResponse, error)
	UpdateEventStatus(eventId, status, remarks string, confirmedDate time.Time) error
}
