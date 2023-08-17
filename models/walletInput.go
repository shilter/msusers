package models

type UpdateWalletInput struct {
	Debet string `json:"Debet"`
    Credit string `json:"Credit"`
	LastBalance string `json:"lastBalance"`
	UserID uint `json:"userID" binding:"required"`
}

type CreateWalletInput struct {
	Debet string `json:"Debet" binding:"required"`
    Credit string `json:"Credit" binding:"required"`
	LastBalance string `json:"lastBalance" binding:"required"`
	UserID uint `json:"userID" binding:"required"`
}