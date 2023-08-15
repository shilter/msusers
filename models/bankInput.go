package models

type UpdateBankInput struct {
	accountBank string `json:"accountBank" binding:"required"`
    nameBank string `json:"nameBank" binding:"required"`
	UserID uint `json:"UserID" binding:"required"`
}