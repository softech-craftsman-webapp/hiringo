package crypto

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt"
)

/*
   |--------------------------------------------------------------------------
   | RSA Public Key @Access
   |--------------------------------------------------------------------------
*/
func RsaAccessPublicKey() (*rsa.PublicKey, error) {
	publicKey := []byte(ReadFile("keys/public.pem"))
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)

	return key, err
}
