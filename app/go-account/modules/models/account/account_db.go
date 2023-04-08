package account

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRole string

const (
	ROLE_ADMIN  AccountRole = "admin"
	ROLE_NORMAL AccountRole = "normal"
)

type Account struct {
	ID        uuid.UUID      `json:"id" gorm:"column:id"`
	Username  string         `json:"username" gorm:"column:username"`
	Password  string         `json:"password" gorm:"password"`
	Role      AccountRole    `json:"role" gorm:"column:role"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
