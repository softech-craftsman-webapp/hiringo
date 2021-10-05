package crypto

import (
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt"
)

/*
   |--------------------------------------------------------------------------
   | ECDSA Public Key @Access
   |--------------------------------------------------------------------------
*/
func EcdsaAccessPublicKey() (*ecdsa.PublicKey, error) {
	publicKey := []byte(ReadFile("keys/public.pem"))
	key, err := jwt.ParseECPublicKeyFromPEM(publicKey)

	return key, err
}
