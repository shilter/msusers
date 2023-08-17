package models

type UpdateBankInput struct {
	AccountBank string `json:"accountBank" binding:"required"`
    NameBank string `json:"nameBank" binding:"required"`
	UserID uint `json:"UserID" binding:"required"`
}

type CreateBankInput struct {
	AccountBank string `json:"accountBank"`
    NameBank string `json:"nameBank"`
	UserID uint `json:"UserID"`
}