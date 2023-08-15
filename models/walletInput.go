package models

type UpdateWalletInput struct {
	Debet string `json:"username" binding:"required"`
    Credit string `json:"password" binding:"required"`
	lastBalance string `json:"lastBalance" binding:"required"`
	UserID uint `json:"userID" binding:"required"`
}