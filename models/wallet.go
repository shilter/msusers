package models

import (
	"msusers/database"
	"time"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Debet string `gorm:"size:255;not null json:"Debet"`
	Credit string `gorm:"size:255;not null" json:"Credit"`
	LastBalance string `gorm:"size:255;not null" json:"lastBalance"`
	UserID uint
	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"createdAt"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updatedAt"`
}

func (wallet *Wallet) Save() (*Wallet, error) {
	err := database.Database.Create(&wallet).Error
	if err != nil {
		return &Wallet{},err
	}
	return wallet, nil
}

func FindWalletsById(id uint) (Wallet, error) {
	var wallet Wallet
	err := database.Database.Where("ID=?", id).Find(&wallet).Error
	if err != nil {
		return Wallet{}, err
	}
	return wallet, nil
}