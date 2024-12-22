package event

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/middleware"
	"event-system-backend/pkg/model/dto"
	"event-system-backend/pkg/model/dto/request"
	eventService "event-system-backend/pkg/service/event"
	"event-system-backend/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventservice eventService.EventService
}

func NewEventController(eventservice eventService.EventService) *EventController {
	return &EventController{eventservice: eventservice}
}

func SetupEventRoutes(r *gin.RouterGroup, ec *EventController) {
	eventGroup := r.Group("/event")
	eventGroup.Use(middleware.AuthMiddleware)
	eventGroup.POST("", ec.CreateEvent)
	eventGroup.GET("/involved", ec.ShowEventsByUserInvolved)
}

func (ec *EventController) CreateEvent(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.Error(handler.NewError(http.StatusInternalServerError, "Failed to retrieve username own user"))
		return
	}

	var body request.CreateEventDto
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

	errEvent := ec.eventservice.CreateEvent(body, username.(string))
	if errEvent != nil {
		c.Error(handler.NewError(errEvent.Code, errEvent.Message))
		return
	}

	c.JSON(http.StatusCreated, dto.BaseResponse(true, "OK", nil))
}

func (ec *EventController) ShowEventsByUserInvolved(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")
	nameStr := c.DefaultQuery("name", "")
	statusStr := c.DefaultQuery("status", "")

	username, exists := c.Get("username")
	if !exists {
		c.Error(handler.NewError(http.StatusInternalServerError, "Failed to retrieve username own user"))
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.Error(handler.NewError(http.StatusBadRequest, "Invalid page parameter"))
		return
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.Error(handler.NewError(http.StatusBadRequest, "Invalid size parameter"))
		return
	}

	events, errEvents := ec.eventservice.ShowEventsByUserInvolved(username.(string), page, size, nameStr, statusStr)
	if errEvents != nil {
		c.Error(handler.NewError(errEvents.Code, errEvents.Message))
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse(true, "OK", events))
}
