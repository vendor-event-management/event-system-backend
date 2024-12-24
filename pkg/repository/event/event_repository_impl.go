package repository

import (
	"event-system-backend/pkg/model/domain"
	"event-system-backend/pkg/model/dto/response"
	"event-system-backend/pkg/utils"
	"time"

	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &EventRepositoryImpl{db: db}
}

func (e *EventRepositoryImpl) CreateEvent(event domain.Event, eventApproval domain.EventApproval) error {
	tx := e.db.Begin()
	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		return err
	}

	eventApproval.EventID = event.ID
	if err := tx.Create(&eventApproval).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (e *EventRepositoryImpl) FindEventByID(eventID string) (domain.Event, error) {
	var response domain.Event
	if err := e.db.Table("events").Where("id = ? AND deleted_at IS NULL", eventID).First(&response).Error; err != nil {
		return response, err
	}
	return response, nil
}

func (e *EventRepositoryImpl) UpdateEventStatus(eventId, status, remarks string, confirmedDate time.Time) error {
	query := e.db.Table("event_approvals").Where("event_approvals.event_id = ?", eventId)
	updatedEvent := domain.EventApproval{Status: domain.EventStatus(status)}
	if !utils.IsEmptyString(remarks) {
		updatedEvent.Remarks = utils.ConvertStringToSQLNullString(remarks)
	}
	if !confirmedDate.IsZero() {
		updatedEvent.ConfirmedDate = utils.ConvertTimeToSQLNullTime(confirmedDate)
	}

	if err := query.Updates(updatedEvent).Error; err != nil {
		return err
	}
	return nil
}

func (e *EventRepositoryImpl) FindAllEventsByUserInvolved(userId string, role domain.RoleType, page, size int, nameEvent, status string) ([]response.EventScanResponse, int64, error) {
	response := []response.EventScanResponse{}

	selectStatement := "events.id AS EventId, events.name AS EventName, users.full_name AS VendorName, events.proposed_dates AS EventProposedDates, " +
		"event_approvals.confirmed_date AS EventConfirmedDate, event_approvals.status AS EventStatus, events.created_at AS CreatedAt"

	baseQuery := e.db.Table("events").Select(selectStatement).
		Joins("JOIN event_approvals ON events.id = event_approvals.event_id").
		Joins("JOIN users ON users.id = event_approvals.vendor_id").
		Where("events.deleted_at IS NULL AND event_approvals.deleted_at IS NULL")

	// filter by role
	if role == domain.HR {
		baseQuery = baseQuery.Where("events.created_by = ?", userId)
	} else if role == domain.Vendor {
		baseQuery = baseQuery.Where("users.id = ?", userId)
	}

	// filter by name event
	if !utils.IsEmptyString(nameEvent) {
		baseQuery = baseQuery.Where("LOWER(events.name) LIKE LOWER(?)", "%"+nameEvent+"%")
	}

	// filter by event status
	if !utils.IsEmptyString(status) {
		baseQuery = baseQuery.Where("event_approvals.status = ?", status)
	}

	// Query for count total data
	var total int64
	if err := baseQuery.Count(&total).Error; err != nil {
		return response, total, err
	}

	// Query with configure paginating event
	offset := (page - 1) * size
	paginatedQuery := baseQuery.Limit(size).Offset(offset).Order("events.created_at DESC")
	if err := paginatedQuery.Scan(&response).Error; err != nil {
		return response, total, err
	}

	return response, total, nil
}

func (e *EventRepositoryImpl) FindDetailEventByID(id string) (response.EventDetailScanResponse, error) {
	response := response.EventDetailScanResponse{}

	selectStatement := "events.id AS EventId, events.name AS EventName, users.full_name AS VendorName, events.proposed_dates AS EventProposedDates, " +
		"event_approvals.confirmed_date AS EventConfirmedDate, event_approvals.status AS EventStatus, events.created_at AS CreatedAt, event_approvals.remarks AS Remarks, events.postal_code AS EventPostalCode, events.location AS EventLocation"

	baseQuery := e.db.Table("events").Select(selectStatement).
		Joins("JOIN event_approvals ON events.id = event_approvals.event_id").
		Joins("JOIN users ON users.id = event_approvals.vendor_id").
		Where("events.id = ? AND events.deleted_at IS NULL", id)

	if err := baseQuery.Scan(&response).Error; err != nil {
		return response, err
	}

	return response, nil
}
