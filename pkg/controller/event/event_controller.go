package event

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/model/dto"
	eventService "event-system-backend/pkg/service/event"
	"event-system-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventservice eventService.EventService
}

func NewEventController(eventservice eventService.EventService) *EventController {
	return &EventController{eventservice: eventservice}
}

func SetupEventRoutes(r *gin.RouterGroup, ec *EventController) {
	healthGroup := r.Group("/event")
	healthGroup.POST("", ec.CreateEvent)
}

func (ec *EventController) CreateEvent(c *gin.Context) {
	var body dto.CreateEventDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(handler.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	if utils.IsEmptyString(body.Name) {
		c.Error(handler.NewError(http.StatusBadRequest, "Event proposed dates must be filled"))
		return
	}

	if utils.IsEmptyString(body.PostalCode) {
		c.Error(handler.NewError(http.StatusBadRequest, "Event location postal code must be filled"))
		return
	}

	if utils.IsEmptyString(body.VendorId) {
		c.Error(handler.NewError(http.StatusBadRequest, "Vendor must be filled"))
		return
	}

	if len(body.ProposedDates) < 1 {
		c.Error(handler.NewError(http.StatusBadRequest, "Event proposed dates must be filled"))
		return
	}

	errEvent := ec.eventservice.CreateEvent(body)
	if errEvent != nil {
		c.Error(handler.NewError(errEvent.Code, errEvent.Message))
		return
	}

	c.JSON(http.StatusCreated, dto.BaseResponse(true, "OK", nil))
}
