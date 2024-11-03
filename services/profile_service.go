package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"dating-app/config"
	"dating-app/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// Max number of swipes allowed per day
const MaxDailySwipes = 3

func GetDailyProfiles(userID uint) ([]models.User, error) {
	var profiles []models.User

	// Fetch up to 10 profiles the user has not seen today
	err := config.DB.Where("id != ?", userID).Limit(10).Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func SwipeProfile(userID, profileID uint, like bool, isPremium bool) error {

	ctx := context.Background()
	swipesKey := fmt.Sprintf("user:%d:swipes", userID)

	if !isPremium {
		// Get current swipe count from Redis
		swipeCount, err := config.RDB.Get(ctx, swipesKey).Int()
		if err != nil && err != redis.Nil {
			return errors.New("error checking daily swipe count")
		}

		// Check daily swipe limit
		if swipeCount >= MaxDailySwipes {
			return errors.New("daily swipe limit reached")
		}
	}

	// Check if profile exist
	var profile models.User
	if err := config.DB.First(&profile, profileID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("profile not found")
		}
		return err
	}

	swipe := models.Swipe{
		UserID:    userID,
		ProfileID: profileID,
		CreatedAt: time.Now(),
		Direction: "pass",
	}

	if like {
		swipe.Direction = "like"
	}

	// Check Existing Swipe Data
	var existingSwipe models.Swipe
	if err := config.DB.Where("user_id = ? AND profile_id = ? AND created_at >= ?", swipe.UserID, swipe.ProfileID, time.Now().Format("2006-01-02")).First(&existingSwipe).Error; err == nil {
		return errors.New("you have already swiped this profile today")
	}

	// Insert Swipe Data
	swipe.CreatedAt = time.Now()
	if err := config.DB.Create(&swipe).Error; err != nil {
		return err
	}

	if !isPremium {
		// Increment swipe count in Redis and Set an expiration for the key if it's a new day
		config.RDB.Incr(ctx, swipesKey)
		config.RDB.Expire(ctx, swipesKey, 24*time.Hour)
	}

	return nil
}

func UpgradePremium(userID uint) error {

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	user.Premium = true
	if err := config.DB.Save(&user).Error; err != nil {
		return errors.New("failed to update user to premium")
	}

	return nil
}
