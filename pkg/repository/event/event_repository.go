package repository

import (
	"event-system-backend/pkg/model/domain"
	"event-system-backend/pkg/model/dto/response"
)

type EventRepository interface {
	CreateEvent(event domain.Event, eventApproval domain.EventApproval) error
	FindAllEventsByUserInvolved(userId string, role domain.RoleType, page, size int, nameEvent, status string) ([]response.EventScanResponse, int64, error)
}
