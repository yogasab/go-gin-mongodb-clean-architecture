package dto

type CheckUserAvailabilityInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CreateNewUserInput struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Location   string `json:"location" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
}

type UpdateUserInput struct {
	ID         string
	Name       string `json:"name,omitempty" binding:"required"`
	Email      string `json:"email,omitempty" binding:"required,email"`
	Password   string `json:"password,omitempty" binding:"required"`
	Location   string `json:"location,omitempty" binding:"required"`
	Occupation string `json:"occupation,omitempty" binding:"required"`
}

type UpdloadUserAvatarInput struct {
	ID     string
	Avatar string `form:"avatar" binding:"required"`
}
