package accountactivity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActivityType string

const (
	ACTIVITY_LOGIN  ActivityType = "login"
	ACTIVITY_LOGOUT ActivityType = "logout"
)

type AccountActivity struct {
	ID        uuid.UUID      `json:"id" gorm:"column:id"`
	UserID    uuid.UUID      `json:"user_id" gorm:"column:user_id"`
	Type      ActivityType   `json:"type" gorm:"column:type"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
