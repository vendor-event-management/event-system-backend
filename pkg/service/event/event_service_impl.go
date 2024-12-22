package event

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/domain"
	"event-system-backend/pkg/model/dto"
	"event-system-backend/pkg/model/dto/request"
	"event-system-backend/pkg/model/dto/response"
	eventrepository "event-system-backend/pkg/repository/event"
	"event-system-backend/pkg/service/user"
	"event-system-backend/pkg/utils"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type EventServiceImpl struct {
	userService     user.UserService
	eventRepository eventrepository.EventRepository
}

func NewEventService(userService user.UserService, eventRepository eventrepository.EventRepository) EventService {
	return &EventServiceImpl{userService: userService, eventRepository: eventRepository}
}

func (e *EventServiceImpl) CreateEvent(data request.CreateEventDto, createdByUser string) *handler.CustomError {
	user, errUser := e.userService.GetUserByIdOrUsernameOrEmail(createdByUser)
	if errUser != nil {
		return handler.NewError(errUser.Code, errUser.Message)
	}

	proposedDates, errProposedDates := utils.ConvertToJSONString(data.ProposedDates)
	if errProposedDates != nil {
		return handler.NewError(http.StatusInternalServerError, errProposedDates.Error())
	}

	event := domain.Event{
		Name:          data.Name,
		PostalCode:    data.PostalCode,
		Location:      utils.ConvertToNullString(data.Location),
		ProposedDates: proposedDates,
		CreatedBy:     user.ID,
	}

	vendor, errVendor := e.userService.GetVendorById(data.VendorId)
	if errVendor != nil {
		return handler.NewError(errVendor.Code, errVendor.Message)
	}

	eventApproval := domain.EventApproval{
		Status:   domain.Pending,
		VendorID: vendor.ID,
	}

	errEvent := e.eventRepository.CreateEvent(event, eventApproval)
	if errEvent != nil {
		return handler.NewError(http.StatusInternalServerError, errEvent.Error())
	}

	return nil
}

func (e *EventServiceImpl) ShowEventsByUserInvolved(userInvolvedID string, page, size int, nameEvent, status string) (*dto.PaginationResponse, *handler.CustomError) {
	user, errUser := e.userService.GetUserByIdOrUsernameOrEmail(userInvolvedID)
	if errUser != nil {
		return nil, handler.NewError(errUser.Code, errUser.Message)
	}

	events, totalEvents, errEvent := e.eventRepository.FindAllEventsByUserInvolved(user.ID.String(), user.Role, page, size, nameEvent, status)
	if errEvent != nil {
		return nil, handler.NewError(http.StatusInternalServerError, errEvent.Error())
	}

	response, errResponse := response.BuildEventResponseFromEventScan(events)
	if errResponse != nil {
		return nil, handler.NewError(http.StatusInternalServerError, errResponse.Error())
	}

	return dto.NewPaginationResponse(page, size, int(totalEvents), response), nil
}

func (e *EventServiceImpl) GetDetailEventByID(id string) (response.EventDetailResponse, *handler.CustomError) {
	responseEvent := response.EventDetailResponse{}

	event, errEvent := e.eventRepository.FindDetailEventByID(id)
	if errEvent != nil {
		if errEvent == gorm.ErrRecordNotFound {
			return responseEvent, handler.NewError(http.StatusNotFound, "Event not found")
		}
		return responseEvent, handler.NewError(http.StatusInternalServerError, errEvent.Error())
	}

	responseEvent, errResponseEvent := response.BuildEventDetailResponseFromEventScan(event)
	if errResponseEvent != nil {
		return responseEvent, handler.NewError(http.StatusInternalServerError, errResponseEvent.Error())
	}

	return responseEvent, nil
}

func (e *EventServiceImpl) ApproveOrRejectEvent(eventID, usernameVendor string, approvalData request.EventApprovalDto) *handler.CustomError {
	var confirmedDate time.Time

	existingEvent, errEvent := e.eventRepository.FindEventByID(eventID)
	if errEvent != nil {
		if errEvent == gorm.ErrRecordNotFound {
			return handler.NewError(http.StatusNotFound, "Event not found")
		}
		return handler.NewError(http.StatusInternalServerError, errEvent.Error())
	}

	vendor, errVendor := e.userService.GetUserByIdOrUsernameOrEmail(usernameVendor)
	if errVendor != nil {
		return handler.NewError(errVendor.Code, errVendor.Message)
	} else if vendor.Role != domain.Vendor {
		return handler.NewError(http.StatusUnauthorized, "Only Vendor can approve or reject the event")
	}

	if !utils.IsEmptyString(approvalData.ConfirmedDate) {
		date, errDate := utils.ConvertStringToTime(approvalData.ConfirmedDate)
		if errDate != nil {
			return handler.NewError(http.StatusBadRequest, "Invalid confirmed date format")
		}
		confirmedDate = date
	}

	errUpdatedEvent := e.eventRepository.UpdateEventStatus(existingEvent.ID.String(), approvalData.Status, approvalData.Remarks, confirmedDate)
	if errUpdatedEvent != nil {
		return handler.NewError(http.StatusInternalServerError, errUpdatedEvent.Error())
	}

	return nil
}
