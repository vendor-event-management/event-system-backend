package response

import (
	"event-system-backend/pkg/utils"
	"time"
)

type EventResponse struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	VendorName    string     `json:"vendorName"`
	ProposedDates []string   `json:"proposedDates"`
	ConfirmedDate *time.Time `json:"confirmedDate"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"createdAt"`
}

func BuildEventResponseFromEventScan(events []EventScanResponse) ([]EventResponse, error) {
	response := []EventResponse{}
	if len(events) > 0 {
		for _, item := range events {
			proposedDates, errDate := utils.ParseDates(item.EventProposedDates)
			if errDate != nil {
				return response, errDate
			}

			response = append(response, EventResponse{
				Id:            item.EventId,
				Name:          item.EventName,
				VendorName:    item.VendorName,
				ProposedDates: proposedDates,
				ConfirmedDate: item.EventConfirmedDate,
				Status:        item.EventStatus,
				CreatedAt:     item.CreatedAt,
			})
		}
	}

	return response, nil
}
