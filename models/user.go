package models

import (
	"msusers/database"
	"time"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"Username"`
	Password string `gorm:"size:255;not null;" json:"Password"`
	Banks []Bank `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Wallets []Wallet `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"createdAt"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updatedAt"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error

	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?",username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Banks").Preload("Wallets").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}