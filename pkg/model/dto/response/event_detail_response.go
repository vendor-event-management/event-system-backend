package response

import (
	"event-system-backend/pkg/utils"
	"time"
)

type EventDetailResponse struct {
	Id                 string     `json:"id"`
	Name               string     `json:"name"`
	VendorName         string     `json:"vendorName"`
	ProposedDates      []string   `json:"proposedDates"`
	ConfirmedDate      *time.Time `json:"confirmedDate"`
	Status             string     `json:"status"`
	Remarks            *string    `json:"remarks"`
	LocationPostalCode string     `json:"locationPostalCode"`
	FullLocation       *string    `json:"fullLocation"`
	CreatedAt          time.Time  `json:"createdAt"`
}

func BuildEventDetailResponseFromEventScan(event EventDetailScanResponse) (EventDetailResponse, error) {
	proposedDates, errDate := utils.ParseDates(event.EventProposedDates)
	if errDate != nil {
		return EventDetailResponse{}, errDate
	}

	return EventDetailResponse{
		Id:                 event.EventId,
		Name:               event.EventName,
		VendorName:         event.VendorName,
		ProposedDates:      proposedDates,
		ConfirmedDate:      event.EventConfirmedDate,
		Status:             event.EventStatus,
		Remarks:            event.Remarks,
		LocationPostalCode: event.EventPostalCode,
		FullLocation:       event.EventLocation,
		CreatedAt:          event.CreatedAt,
	}, nil
}
