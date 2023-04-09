package jwt

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

const rsaThumbprintTemplate = `{"e":"%s","kty":"RSA","n":"%s"}`

type Jwk struct {
	Alg string `json:"alg"`
	Kty string `json:"kty"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
	Kid string `json:"kid"`
}

type Jwks struct {
	Keys []Jwk `json:"keys"`
}

func GenJwt() (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": "example.com",
		"iat": now.Unix(),
		"foo": "bar",
		"exp": now.Add(time.Hour * 24 * 365).Unix(),
		"example.com": map[string]string{
			"nested": "nested-value",
		},
	})

	key, err := os.ReadFile("./keys/jwt.key.pem")
	if err != nil {
		return "", err
	}

	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return "", err
	}

	return token.SignedString(rsaKey)
}

func GenJwks() (string, error) {
	jwk := Jwk{Alg: "RS256", Kty: "RSA", Use: "sig"}

	key, err := os.ReadFile("./keys/jwt.key.pem")
	if err != nil {
		panic(err)
	}

	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return "", err
	}
	pubKey := rsaKey.PublicKey

	jwk.N = base64.RawURLEncoding.EncodeToString(pubKey.N.Bytes())

	d := make([]byte, 8)
	binary.BigEndian.PutUint64(d, uint64(pubKey.E))
	jwk.E = base64.RawURLEncoding.EncodeToString(bytes.TrimLeft(d, "\x00"))

	template := fmt.Sprintf(rsaThumbprintTemplate, jwk.N, jwk.E)
	h := md5.New()
	_, _ = h.Write([]byte(template))
	jwk.Kid = hex.EncodeToString(h.Sum(nil))

	jwks := Jwks{
		Keys: []Jwk{jwk},
	}
	js, err := json.Marshal(jwks)
	if err != nil {
		return "", err
	}

	return string(js), nil
}
