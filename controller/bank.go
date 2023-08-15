package controller

import (
	"msusers/database"
	"msusers/models"
	"msusers/helper"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SaveBank(context *gin.Context) {
	var input models.Bank

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

	savedBank, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data":savedBank})
}

func ListBank(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Banks})
}

func ListAllBank(context *gin.Context) {
	var bank []models.Bank
	database.Database.Find(&bank)
	context.JSON(http.StatusOK, gin.H{"data": bank})
}

func GetBankDetailsById(context *gin.Context) {
	var bank models.Bank
	bankId := context.Param("id")
	if err := database.Database.Where("id = ?", bankId).First(&bank).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"data": bank})
}

func UpdatedBank(context *gin.Context) {
	var bank models.Bank
	bankId := context.Param("id")
	if err := database.Database.Where("id = ?", bankId).First(&bank).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	var input models.UpdateBankInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.Model(&bank).Updates(input)
	context.JSON(http.StatusOK, gin.H{"data": bank})
}

func DeleteBank(context *gin.Context) {
	var bank models.Bank
	bankId := context.Param("id")
  	if err := database.Database.Where("id = ?", bankId).First(&bank).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    	return
  	}

  	database.Database.Delete(&bank)
	context.JSON(http.StatusOK, gin.H{"data": true})
}