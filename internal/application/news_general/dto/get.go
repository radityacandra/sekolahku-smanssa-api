package dto

type DetailNewsRequest struct {
	ID int `param:"id" validate:"required"`
}

type DetailNewsResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	PublishedAt string `json:"publishedAt"`
	Author      string `json:"author"`
}
