package dto

type CheckUserAvailabilityInput struct {
	Email string `json:"email" binding:"required,email"`
}
