package request

type CreateEventDto struct {
	Name          string   `json:"name"`
	PostalCode    string   `json:"postalCode"`
	Location      *string  `json:"location"`
	ProposedDates []string `json:"proposedDates"`
	VendorId      string   `json:"vendorId"`
}
