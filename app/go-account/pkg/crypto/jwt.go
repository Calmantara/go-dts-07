package crypto

import (
	"errors"

	"github.com/kataras/jwt"
)

const sharedKey = ("sercrethatmaycontainch@r$32chars")

func SignJWT(claim any) (token string, err error) {
	// sign jwt
	tkn, err := jwt.Sign(jwt.HS256, []byte(sharedKey), claim)
	if err != nil {
		err = errors.New("error sign claim")
	}
	return string(tkn), err
}

func ParseJWT(token string, claims any) (err error) {
	// Verify and extract claims from a token:
	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(sharedKey), []byte(token))
	if err != nil {
		err = errors.New("error parse token")
		return
	}
	err = verifiedToken.Claims(&claims)
	return err
}
