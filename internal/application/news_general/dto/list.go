package dto

type NewsGeneralRequest struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}

type NewsGeneralResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
