package handlers

import (
	"net/http"

	"github.com/MTursynbekov/GolangApplicationDevelopment/hw7/models"

	"github.com/gin-gonic/gin"
)

type CreateAdInput struct {
	Title       string   `json:"title" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Images      []string `json:"images" binding:"required"`
	Price       int      `json:"price" binding:"required"`
}

type UpdateAdInput struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Price       int      `json:"price"`
}

func GetAllAds(context *gin.Context) {
	var ads []models.Ad
	err := models.DB.Find(&ads).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": ads})
}

func CreateAd(context *gin.Context) {

	var input CreateAdInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ad := models.Ad{
		Title:       input.Title,
		Author:      input.Author,
		Description: input.Description,
		Images:      input.Images,
		Price:       input.Price,
	}
	err := models.DB.Create(&ad).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": ad})
}

func UpdateAd(context *gin.Context) {
	var ad models.Ad

	if err := models.DB.Where("id=?", context.Param("id")).First(&ad).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateAdInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DB.Model(&ad).Update(input).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": ad})
}

func DeleteAd(context *gin.Context) {
	var ad models.Ad

	if err := models.DB.Where("id=?", context.Param("id")).First(&ad).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	err := models.DB.Delete(&ad).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": "deleted: true"})
}
