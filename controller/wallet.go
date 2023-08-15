package controller

import (
	"msusers/database"
	"msusers/models"
	"msusers/helper"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SaveWallet(context *gin.Context) {
	var input models.Wallet

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	User, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	input.UserID = User.ID

	savedWallets, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data":savedWallets})
}

func ListWallet(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Banks})
} 

func ListAllWallets(context *gin.Context) {
	var wallets []models.Wallet
	database.Database.Find(&wallets)
	context.JSON(http.StatusOK, gin.H{"data": wallets})
}
	
func GetWalletsById(context *gin.Context) {
	var wallet models.Wallet
	walletId := context.Param("id")
	if err := database.Database.Where("id = ?", walletId).First(&wallet).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"data": wallet})
}
	
func UpdatedWallet(context *gin.Context) {
	var wallet models.Wallet
	walletId := context.Param("id")
	if err := database.Database.Where("id = ?", walletId).First(&wallet).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	var input models.UpdateWalletInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.Model(&wallet).Updates(input)
	context.JSON(http.StatusOK, gin.H{"data": wallet})
}
	
func DeleteWallet(context *gin.Context) {
	var wallet models.Wallet
	walletId := context.Param("id")
  	if err := database.Database.Where("id = ?", walletId).First(&wallet).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    	return
  	}

  	database.Database.Delete(&wallet)
	context.JSON(http.StatusOK, gin.H{"data": true})
}