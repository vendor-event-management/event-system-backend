package event

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto"
	"event-system-backend/pkg/model/dto/request"
)

type EventService interface {
	CreateEvent(data request.CreateEventDto, createdByUser string) *handler.CustomError
	ShowEventsByUserInvolved(userInvolved string, page, size int, nameEvent, status string) (*dto.PaginationResponse, *handler.CustomError)
}
