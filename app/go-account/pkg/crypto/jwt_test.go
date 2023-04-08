package crypto

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/Calmantara/go-account/modules/models/token"
	"github.com/stretchr/testify/assert"
)

func TestSignJWT(t *testing.T) {
	type (
		input struct {
			claim any
		}
		want struct {
			token string
			err   error
		}
	)

	testCases := []struct {
		desc  string
		input input
		want  want
	}{
		{
			desc: "happy case",
			input: input{ // inputan dari function kita
				claim: token.DefaultClaim{
					Expired:   1234567899999999,
					NotBefore: 12345678,
					IssuedAt:  12345678,
					Issuer:    "http://go-account",
					Audience:  "http://dts-07",
					JTI:       "this-is-jti",
					Type:      token.ID_TOKEN,
				},
			},
			want: want{ // output yang kita inginkan
				err:   nil,
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEyMzQ1Njc4OTk5OTk5OTksIm5iZiI6MTIzNDU2NzgsImlhdCI6MTIzNDU2NzgsImlzcyI6Imh0dHA6Ly9nby1hY2NvdW50IiwiYXVkIjoiaHR0cDovL2R0cy0wNyIsImp0aSI6InRoaXMtaXMtanRpIiwidHlwIjoiaWRfdG9rZW4ifQ.7w0E_SGKyS56D8AjwVSGeMcuBN064uxCYX1wHT53Q1w",
			},
		},
		{
			desc: "error sign",
			input: input{ // inputan dari function kita
				claim: make(chan int),
			},
			want: want{ // output yang kita inginkan
				err:   errors.New("error sign claim"),
				token: "",
			},
		},
	}
	// loop over test cases
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tkn, err := SignJWT(tC.input.claim) // actual yang kita dapatkan dari function

			if tC.want.err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tC.want.token, tkn)
		})
	}
}

func TestParseJwt(t *testing.T) {
	type (
		input struct {
			token  string
			claims any
		}
		want struct {
			err    error
			claims any
		}
	)

	testCases := []struct {
		desc  string
		input input
		want  want
	}{
		{
			desc: "happy case",
			input: input{
				token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEyMzQ1Njc4OTk5OTk5OTksIm5iZiI6MTIzNDU2NzgsImlhdCI6MTIzNDU2NzgsImlzcyI6Imh0dHA6Ly9nby1hY2NvdW50IiwiYXVkIjoiaHR0cDovL2R0cy0wNyIsImp0aSI6InRoaXMtaXMtanRpIiwidHlwIjoiaWRfdG9rZW4ifQ.7w0E_SGKyS56D8AjwVSGeMcuBN064uxCYX1wHT53Q1w",
				claims: token.DefaultClaim{},
			},
			want: want{
				err: nil,
				claims: token.DefaultClaim{
					Expired:   1234567899999999,
					NotBefore: 12345678,
					IssuedAt:  12345678,
					Issuer:    "http://go-account",
					Audience:  "http://dts-07",
					JTI:       "this-is-jti",
					Type:      token.ID_TOKEN,
				},
			},
		},
		{
			desc: "happy case",
			input: input{
				token:  "",
				claims: token.DefaultClaim{},
			},
			want: want{
				err:    errors.New("error parse token"),
				claims: token.DefaultClaim{},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := ParseJWT(tC.input.token, &tC.input.claims)
			if tC.want.err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
			}

			var act token.DefaultClaim
			b, _ := json.Marshal(tC.input.claims)
			json.Unmarshal(b, &act)

			assert.EqualValues(t, tC.want.claims, act)
		})
	}
}
