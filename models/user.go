package models

type User struct {
	ID int `json:"id,omitempty"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
