package dto

type HeadlineListRequest struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}

type HeadlineListResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
