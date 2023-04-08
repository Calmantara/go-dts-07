package crypto

import "golang.org/x/crypto/bcrypt"

func GenerateHash(payload string) (hashed string, err error) {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	return string(hashedPassword), err
}

func CompareHash(hashed, password string) (err error) {
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return
}
