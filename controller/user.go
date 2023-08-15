package controller

import (
	"msusers/database"
	"msusers/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var input models.CreateUserInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()


	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func ListAllUser(context *gin.Context) {
	var user []models.User
	database.Database.Find(&user)
	context.JSON(http.StatusOK, gin.H{"data": user})
} 
	
func GetUserDetails(context *gin.Context) {
	var user models.User
	userId := context.Param("id")
	if err := database.Database.Where("id = ?", userId).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"data": user})
}
	
func UpdatedUser(context *gin.Context) {
	var user models.User
	userId := context.Param("id")
	if err := database.Database.Where("id = ?", userId).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	var input models.UpdateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.Model(&user).Updates(input)
	context.JSON(http.StatusOK, gin.H{"data": user})
}
	
func DeleteUser(context *gin.Context) {
	var user models.User
	userId := context.Param("id")
  	if err := database.Database.Where("id = ?", userId).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    	return
  	}

  	database.Database.Delete(&user)
	context.JSON(http.StatusOK, gin.H{"data": true})
}