package types

import (
	"errors"
	"fmt"
)

type Credentials struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
}

type Error struct {
	Error   *string `json:"error"`
	Message *string `json:"message"`
}

type AccessTokenResponse struct {
	Error
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	Scope       string `json:"scope"`
}

func GetError(e Error) (err error) {
	if e.Message != nil || e.Error != nil {
		err = errors.New(fmt.Sprintf("%s | error code: %s", *e.Message, *e.Error))

		return
	}

	return
}
