package dto

type DetailProfileRequest struct {
	ID int `param:"id" validate:"required"`
}
