package token

type TokenType string

const (
	ID_TOKEN      TokenType = "id_token"
	ACCESS_TOKEN  TokenType = "access_token"
	REFRESH_TOKEN TokenType = "refresh_token"
)

type Tokens struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type DefaultClaim struct {
	// exp	kapan token expired
	// nbf	kapan token boleh digunakan
	// iat	kapan token dibuat
	// iss	siapa yang membuat token
	// aud	siapa yang akan menggunakan token
	// jti	unique identifier (id account_activity)
	// typ	id_token / access_token / refresh_token
	Expired   int       `json:"exp"`
	NotBefore int       `json:"nbf"`
	IssuedAt  int       `json:"iat"`
	Issuer    string    `json:"iss"`
	Audience  string    `json:"aud"`
	JTI       string    `json:"jti"`
	Type      TokenType `json:"typ"`
}

type IDClaim struct {
	// name
	// family_name
	// given_name
	// middle_name
	// nickname
	// preferred_username
	// profile
	// picture
	// website
	// gender
	// birthdate
	// zoneinfo
	// locale
	// updated_at
	Username string `json:"preferred_username"`
	Role     string `json:"role"`
}

type AccessClaim struct {
	Role   string `json:"role"`
	UserID string `json:"user_id"`
}
