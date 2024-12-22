package dto

type PaginationData struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

type PaginationResponse struct {
	Content    interface{}    `json:"content"`
	Pagination PaginationData `json:"pagination"`
}

func NewPaginationResponse(page, size, total int, data interface{}) *PaginationResponse {
	totalPages := (total + size - 1) / size

	return &PaginationResponse{
		Content: data,
		Pagination: PaginationData{
			Page:       page,
			Size:       size,
			Total:      total,
			TotalPages: totalPages,
		},
	}
}

func (p *PaginationResponse) Offset() int {
	return (p.Pagination.Page - 1) * p.Pagination.Size
}
