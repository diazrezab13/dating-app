package controllers

import (
	"dating-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfiles(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	profiles, err := services.GetDailyProfiles(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profiles": profiles})
}

func Swipe(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	isPremium := c.GetBool("status")

	var input struct {
		ProfileID uint `json:"profile_id" binding:"required"`
		Like      bool `json:"like"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.SwipeProfile(userID, input.ProfileID, input.Like, isPremium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Swipe registered"})
}

func UpgradePremium(c *gin.Context) {

	var input struct {
		UserID uint `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpgradePremium(input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "account upgraded"})
}
