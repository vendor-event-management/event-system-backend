package response

import "time"

type EventDetailScanResponse struct {
	EventId            string
	EventName          string
	VendorName         string
	EventProposedDates string
	EventConfirmedDate *time.Time
	EventStatus        string
	CreatedAt          time.Time
	Remarks            *string
	EventPostalCode    string
	EventLocation      *string
}
