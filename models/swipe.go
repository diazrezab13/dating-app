package models

import "time"

type Swipe struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	ProfileID uint      `gorm:"not null"`
	Direction string    `gorm:"type:varchar(4);check:Direction IN ('like', 'pass')"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
