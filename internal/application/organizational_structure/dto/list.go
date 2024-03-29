package dto

type OrganizationStructureRequest struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}

type OrganizationStructureResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Title string `json:"title"`
}
