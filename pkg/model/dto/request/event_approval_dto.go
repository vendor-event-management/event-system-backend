package request

type EventApprovalDto struct {
	Status        string `json:"status"`
	Remarks       string `json:"remarks"`
	ConfirmedDate string `json:"confirmedDate"`
}
