package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var KeyFunc = func(t *jwt.Token) (interface{}, error) {
	keyData, err := os.ReadFile("./keys/jwt.pub.pem")
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return "", err
	}
	return key, nil
}

func TestGenJwt(t *testing.T) {
	got, err := GenJwt()
	if err != nil {
		t.Errorf("want no err, but got: %s", err.Error())
	}

	// verify jwt
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(got, claims, KeyFunc)

	assert.Nil(t, err)
	assert.Equal(t, true, token.Valid)
}

func TestGenJwks(t *testing.T) {
	jwks, err := GenJwks()
	assert.Nil(t, err)
	fmt.Println(jwks)
}
