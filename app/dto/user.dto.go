package dto

type CheckUserAvailabilityInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CreateNewUserInput struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	AvatarFileName string `json:"avatar" binding:"required"`
	Location       string `json:"location" binding:"required"`
	Occupation     string `json:"occupation" binding:"required"`
}
