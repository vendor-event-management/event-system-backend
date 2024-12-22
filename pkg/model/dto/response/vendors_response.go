package response

import "event-system-backend/pkg/model/domain"

type VendorsResponse struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
}

func BuildVendorsResponseFromFetchedVendors(data []domain.User) []VendorsResponse {
	response := []VendorsResponse{}
	if len(data) > 0 {
		for _, vendor := range data {
			response = append(response, VendorsResponse{
				ID:       vendor.ID.String(),
				FullName: vendor.FullName,
			})
		}
	}
	return response
}
