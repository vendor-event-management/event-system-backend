package repository

import "event-system-backend/pkg/model/domain"

type EventRepository interface {
	CreateEvent(event domain.Event, eventApproval domain.EventApproval) error
}
