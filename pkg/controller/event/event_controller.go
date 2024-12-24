package event

import (
	"event-system-backend/pkg/handler"
	"event-system-backend/pkg/middleware"
	"event-system-backend/pkg/model/domain"
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
	eventGroup.GET("/:eventId", ec.GetDetailEventByID)
	eventGroup.GET("/by-user/:userId", ec.ShowEventsByUserInvolved)
	eventGroup.PUT("/:eventId/approval", ec.UpdateEventApprovalStatus)
}

// CreateEvent handles the creation of an event
// @Summary Create a new event
// @Description Create a new event with the provided details
// @Tags Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <your-token-here>"
// @Param body body request.CreateEventDto true "Event data. Propose date format should be ['dd-mm-yyyy', 'dd-mm-yyyy']"
// @Success 201 {object} dto.Response "Event created successfully"
// @Failure 400 {object} dto.Response "Bad request: Missing required fields"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /event [post]
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

// ShowEventsByUserInvolved shows events by user involvement
// @Summary Show events by user involvement
// @Description Get the list of events that a user is involved in, with pagination
// @Tags Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <your-token-here>"
// @Param userId path string true "User ID"
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Param name query string false "Event name filter"
// @Param status query string false "Event status filter"
// @Success 200 {object} dto.Response{data=dto.PaginationResponse{content=[]response.EventResponse}} "Events retrieved successfully"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /event/by-user/{userId} [get]
func (ec *EventController) ShowEventsByUserInvolved(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")
	nameStr := c.DefaultQuery("name", "")
	statusStr := c.DefaultQuery("status", "")
	userId := c.Param("userId")

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

	events, errEvents := ec.eventservice.ShowEventsByUserInvolved(userId, page, size, nameStr, statusStr)
	if errEvents != nil {
		c.Error(handler.NewError(errEvents.Code, errEvents.Message))
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse(true, "OK", events))
}

// GetDetailEventByID retrieves event details by event ID
// @Summary Get event details by ID
// @Description Get the details of a specific event by its ID
// @Tags Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <your-token-here>"
// @Param eventId path string true "Event ID"
// @Success 200 {object} dto.Response{data=response.EventDetailResponse} "Event details retrieved successfully"
// @Failure 400 {object} dto.Response "Bad request"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /event/{eventId} [get]
func (ec *EventController) GetDetailEventByID(c *gin.Context) {
	eventId := c.Param("eventId")

	event, err := ec.eventservice.GetDetailEventByID(eventId)
	if err != nil {
		c.Error(handler.NewError(err.Code, err.Message))
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse(true, "OK", event))
}

// UpdateEventApprovalStatus updates the approval status of an event
// @Summary Update the approval status of an event
// @Description Approve or reject an event by changing its status
// @Tags Event
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <your-token-here>"
// @Param eventId path string true "Event ID"
// @Param body body request.EventApprovalDto true "Approval data. Status should be 'Approved' or 'Rejected'"
// @Success 200 {object} dto.Response "Event status updated successfully"
// @Failure 400 {object} dto.Response "Bad request: Invalid status"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /event/{eventId}/approval [put]
func (ec *EventController) UpdateEventApprovalStatus(c *gin.Context) {
	var body request.EventApprovalDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(handler.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	eventId := c.Param("eventId")
	username, exists := c.Get("username")
	if !exists {
		c.Error(handler.NewError(http.StatusInternalServerError, "Failed to retrieve username own user"))
		return
	}

	if utils.IsEmptyString(body.Status) || (body.Status != string(domain.Approved) && body.Status != string(domain.Rejected)) {
		c.Error(handler.NewError(http.StatusBadRequest, "Invalid event status request"))
		return
	}

	errApproval := ec.eventservice.ApproveOrRejectEvent(eventId, username.(string), body)
	if errApproval != nil {
		c.Error(handler.NewError(errApproval.Code, errApproval.Message))
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse(true, "OK", nil))
}
