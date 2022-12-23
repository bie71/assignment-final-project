package http_response

type Result struct {
	Page       any `json:"page"`
	PerPage    any `json:"per_page"`
	TotalCount any `json:"total_count"`
	Records    any `json:"records"`
}

func PaginationInfo(page, perPage, totalCount, records any) *Result {
	return &Result{
		Page:       page,
		PerPage:    perPage,
		TotalCount: totalCount,
		Records:    records,
	}
}
