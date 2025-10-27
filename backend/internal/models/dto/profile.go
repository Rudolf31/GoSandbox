package dto

type ProfileDTO struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
	Id       int32  `json:"id"`
}

type NewProfileDTO struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Age      int    `json:"age" binding:"required,min=18,max=100"`
}

type PatchProfileDTO struct {
	Name     *string `json:"name"`
	LastName *string `json:"last_name"`
	Age      *int    `json:"age" binding:"omitempty,min=18,max=100"`
}
