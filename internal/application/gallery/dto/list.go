package dto

type GalleryRequest struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}

type GalleryResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}
