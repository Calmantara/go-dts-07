package account

type CreateAccount struct {
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Role     AccountRole `json:"role" binding:"required"`
}

type LoginAccount struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
