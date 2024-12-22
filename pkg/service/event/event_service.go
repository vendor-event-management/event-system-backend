package event

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto"
	"event-system-backend/pkg/model/dto/request"
	"event-system-backend/pkg/model/dto/response"
)

type EventService interface {
	CreateEvent(data request.CreateEventDto, createdByUser string) *handler.CustomError
	ShowEventsByUserInvolved(userInvolvedID string, page, size int, nameEvent, status string) (*dto.PaginationResponse, *handler.CustomError)
	GetDetailEventByID(id string) (response.EventDetailResponse, *handler.CustomError)
	ApproveOrRejectEvent(eventID, usernameVendor string, approvalData request.EventApprovalDto) *handler.CustomError
}
