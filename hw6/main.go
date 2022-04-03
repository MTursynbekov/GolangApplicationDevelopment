package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ad struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Price       float64  `json:"price"`
}

var ads = []Ad{
	{
		ID:          "1",
		Title:       "FARMTRAC 45 CLASSIC TRACTOR PRICE AND SPECIFICATIONS",
		Author:      "John",
		Description: "Farmtrac tractor is one of the most well-known tractors brands in India offering several models. Farmtrac 45 is ad budget-friendly tractor in this segment.",
		Images: []string{
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/1/image/154",
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/1/image/156",
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/1/image/158",
		},
		Price: 1500,
	},
	{
		ID:          "2",
		Title:       "BLANKET - BUY BED BLANKETS ONLINE AT BEST PRICES IN INDIA",
		Author:      "Katrine",
		Description: "Blankets at Upto 55% OFF-Explore Latest Collections of Blankets. Shop from ad wide range of Blankets Online such as Double Bed Blankets, Single Bed Blanket, King Size Blankets & more. Great Offers-Free Shipping.",
		Images: []string{
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/2/image/54",
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/2/image/56",
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/2/image/58",
		},
		Price: 990,
	},
	{
		ID:          "3",
		Title:       "CUSTOM CAT SOCKS",
		Author:      "Lois",
		Description: "Our super cute and comfy custom cat socks make the best cat parent gift.Personalized design just for you:  Our sock designers will first create ad personalized design artwork for your approval before proceeding with production!",
		Images: []string{
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/1/image/154",
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/1/image/156",
			"https://example.s3-ap-northeast-1.amazonaws.com/uploads/ad/1/image/158",
		},
		Price: 10,
	},
}

func getAds(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ads)
}

func postAds(c *gin.Context) {
	var ad Ad

	if err := c.BindJSON(&ad); err != nil {
		return
	}

	for _, a := range ads {
		if a.ID == ad.ID {
			c.IndentedJSON(http.StatusBadRequest, "Ad with this ID already exists")
			return
		}
	}
	ads = append(ads, ad)
	c.IndentedJSON(http.StatusCreated, ad)
}

func getAdByID(c *gin.Context) {
	id := c.Param("id")

	for _, ad := range ads {
		if ad.ID == id {
			c.IndentedJSON(http.StatusOK, ad)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ad not found"})
}

func deleteAdByID(c *gin.Context) {
	id := c.Param("id")

	for i, ad := range ads {
		if ad.ID == id {
			ads = append(ads[:i], ads[i+1:]...)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/ads", getAds)
	router.POST("/ads", postAds)
	router.GET("/ads/:id", getAdByID)
	router.DELETE("/ads/:id", deleteAdByID)

	router.Run("localhost:8080")
}
