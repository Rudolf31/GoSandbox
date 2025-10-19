package dto

type ProfileDTO struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      string `json:"age"`
	Id       int    `json:"id"`
}

type NewProfileDTO struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Age      string `json:"age" binding:"required,min=18,max=100"`
}
