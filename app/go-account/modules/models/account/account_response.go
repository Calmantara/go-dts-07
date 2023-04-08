package account

import (
	"time"

	"github.com/google/uuid"
)

type AccountResponse struct {
	ID        uuid.UUID   `json:"id"`
	Username  string      `json:"username"`
	Role      AccountRole `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
}

type AccountResponseWithPassword struct {
	AccountResponse
	Password string
}
