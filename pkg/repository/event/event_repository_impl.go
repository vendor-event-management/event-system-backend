package repository

import (
	"event-system-backend/pkg/model/domain"

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
