package response

import "time"

type EventScanResponse struct {
	EventId            string
	EventName          string
	VendorName         string
	EventProposedDates string
	EventConfirmedDate *time.Time
	EventStatus        string
	CreatedAt          time.Time
}
