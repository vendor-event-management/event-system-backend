package event

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto"
)

type EventService interface {
	CreateEvent(data dto.CreateEventDto, createdByUser string) *handler.CustomError
}
