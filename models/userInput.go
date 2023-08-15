package models

type UpdateUserInput struct {
	Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}