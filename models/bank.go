package models

import (
	"msusers/database"
	"time"
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	accountBank string `gorm:"size:255;not null;unique json:"accountBank"`
	nameBank string `gorm:"size:255;not null" json:"nameBank"`
	UserID uint
	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"createdAt"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updatedAt"`
}

func (bank *Bank) Save() (*Bank, error) {
	err := database.Database.Create(&bank).Error
	if err != nil {
		return &Bank{},err
	}
	return bank, nil
}

func FindBanksById(id uint) (Bank, error) {
	var bank Bank
	err := database.Database.Where("ID=?", id).Find(&bank).Error
	if err != nil {
		return Bank{}, err
	}
	return bank, nil
}